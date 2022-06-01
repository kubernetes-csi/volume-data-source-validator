/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package data_source_validator

import (
	"fmt"

	volumesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1"
	popv1beta1 "github.com/kubernetes-csi/volume-data-source-validator/client/apis/volumepopulator/v1beta1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamiclister"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	"github.com/kubernetes-csi/volume-data-source-validator/pkg/metrics"
)

type populatorController struct {
	dynClient     dynamic.Interface
	client        kubernetes.Interface
	eventRecorder record.EventRecorder
	queue         workqueue.RateLimitingInterface

	popLister       dynamiclister.Lister
	popListerSynced cache.InformerSynced
	pvcLister       corelisters.PersistentVolumeClaimLister
	pvcListerSynced cache.InformerSynced

	metrics metrics.MetricsManager
}

var (
	pvcGK            = metav1.GroupKind{Group: v1.GroupName, Kind: "PersistentVolumeClaim"}
	volumeSnapshotGK = metav1.GroupKind{Group: volumesnapshotv1.GroupName, Kind: "VolumeSnapshot"}

	PopulatorResource = popv1beta1.SchemeGroupVersion.WithResource("volumepopulators")
)

func NewDataSourceValidator(
	dynClient dynamic.Interface,
	client kubernetes.Interface,
	volumePopulatorInformer cache.SharedIndexInformer,
	pvcInformer coreinformers.PersistentVolumeClaimInformer,
	metrics metrics.MetricsManager,
) *populatorController {
	broadcaster := record.NewBroadcaster()
	broadcaster.StartLogging(klog.Infof)
	broadcaster.StartRecordingToSink(&corev1.EventSinkImpl{Interface: client.CoreV1().Events(v1.NamespaceAll)})
	var eventRecorder record.EventRecorder
	eventRecorder = broadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: fmt.Sprintf("volume-data-source-validator")})

	ctrl := &populatorController{
		dynClient:     dynClient,
		client:        client,
		eventRecorder: eventRecorder,
		metrics:       metrics,
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "pvc"),
	}

	pvcInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    ctrl.enqueueWork,
			UpdateFunc: func(oldObj, newObj interface{}) { ctrl.enqueueWork(newObj) },
			DeleteFunc: ctrl.enqueueWork,
		},
	)
	ctrl.pvcLister = pvcInformer.Lister()
	ctrl.pvcListerSynced = pvcInformer.Informer().HasSynced

	ctrl.popLister = dynamiclister.New(volumePopulatorInformer.GetIndexer(), PopulatorResource)
	ctrl.popListerSynced = volumePopulatorInformer.HasSynced

	return ctrl
}

func (ctrl *populatorController) Run(workers int, stopCh <-chan struct{}) {
	defer ctrl.queue.ShutDown()

	klog.Infof("Starting volume-data-source-validator controller")
	defer klog.Infof("Shutting down volume-data-source-validator controller")

	if !cache.WaitForCacheSync(stopCh, ctrl.popListerSynced, ctrl.pvcListerSynced) {
		klog.Errorf("Cannot sync caches")
		return
	}

	for i := 0; i < workers; i++ {
		go wait.Until(ctrl.worker, 0, stopCh)
	}

	<-stopCh
}

// enqueueWork adds PVC to given work queue.
func (ctrl *populatorController) enqueueWork(obj interface{}) {
	// Beware of "xxx deleted" events
	if unknown, ok := obj.(cache.DeletedFinalStateUnknown); ok && unknown.Obj != nil {
		obj = unknown.Obj
	}
	if pvc, ok := obj.(*v1.PersistentVolumeClaim); ok {
		objName, err := cache.DeletionHandlingMetaNamespaceKeyFunc(pvc)
		if err != nil {
			klog.Errorf("failed to get key from object: %v, %v", err, pvc)
			return
		}
		klog.V(5).Infof("enqueued %q for sync", objName)
		ctrl.queue.Add(objName)
	}
}

// worker is the main worker for PVCs.
func (ctrl *populatorController) worker() {
	keyObj, quit := ctrl.queue.Get()
	if quit {
		return
	}
	defer ctrl.queue.Done(keyObj)

	if err := ctrl.syncPvcByKey(keyObj.(string)); err != nil {
		// Rather than wait for a full resync, re-add the key to the
		// queue to be processed.
		ctrl.queue.AddRateLimited(keyObj)
		klog.V(4).Infof("Failed to sync pvc %q, will retry again: %v", keyObj.(string), err)
	} else {
		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		ctrl.queue.Forget(keyObj)
	}
}

// syncPvcByKey processes a PVC request.
func (ctrl *populatorController) syncPvcByKey(key string) error {
	klog.V(5).Infof("syncPvcByKey[%s]", key)

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	klog.V(5).Infof("worker: pvc namespace [%s] name [%s]", namespace, name)
	if err != nil {
		klog.Errorf("error getting namespace & name of pvc %q to get pvc from informer: %v", key, err)
		return nil
	}
	pvc, err := ctrl.pvcLister.PersistentVolumeClaims(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("pvc '%s' in work queue no longer exists", key))
			return nil
		}
		klog.V(2).Infof("error getting pvc %q from informer: %v", key, err)
		return err
	}

	dataSourceRef := pvc.Spec.DataSourceRef
	if dataSourceRef == nil {
		// No data source
		ctrl.metrics.IncrementCount(metrics.DataSourceEmptyResultName)
		return nil
	}
	apiGroup := ""
	if dataSourceRef.APIGroup != nil {
		apiGroup = *dataSourceRef.APIGroup
	}

	gk := metav1.GroupKind{
		Group: apiGroup,
		Kind:  dataSourceRef.Kind,
	}
	klog.V(3).Infof("PVC %q datasource is %q", pvc.Name, gk.String())

	valid, err := ctrl.validateGroupKind(gk)
	if err != nil {
		return err
	}

	if !valid {
		ctrl.eventRecorder.Event(pvc, v1.EventTypeWarning, "UnrecognizedDataSourceKind",
			"The datasource for this PVC does not match any registered VolumePopulator")
	}

	return nil
}

func (ctrl *populatorController) validateGroupKind(gk metav1.GroupKind) (bool, error) {
	// Cloning PVCs and Volume Snapshots are special cases, allowed by the
	// core, so don't reject these.
	switch gk {
	case pvcGK:
		ctrl.metrics.IncrementCount(metrics.DataSourcePVCResultName)
		klog.V(4).Infof("Allowing PVC as a special case")
		return true, nil
	case volumeSnapshotGK:
		ctrl.metrics.IncrementCount(metrics.DataSourceSnapshotResultName)
		klog.V(4).Infof("Allowing VolumeSnapshot as a special case")
		return true, nil
	}
	unstPopulators, err := ctrl.popLister.List(labels.Everything())
	if err != nil {
		klog.Errorf("Failed to list populators: %v", err)
		ctrl.metrics.IncrementCount(metrics.DataSourceErrorResultName)
		return false, err
	}
	for _, unstPopulator := range unstPopulators {
		var populator popv1beta1.VolumePopulator
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstPopulator.UnstructuredContent(), &populator)
		if err != nil {
			ctrl.metrics.IncrementCount(metrics.DataSourceErrorResultName)
			return false, err
		}
		if populator.SourceKind == gk {
			ctrl.metrics.IncrementCount(metrics.DataSourcePopulatorResultName)
			klog.V(4).Infof("Allowing %q due to %q populator", gk.String(), populator.Name)
			return true, nil
		}
	}
	ctrl.metrics.IncrementCount(metrics.DataSourceInvalidResultName)
	klog.Warningf("No populator matches %s", gk.String())
	return false, nil
}
