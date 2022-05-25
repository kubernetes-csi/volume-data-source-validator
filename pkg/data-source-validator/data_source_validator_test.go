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
	"errors"
	"net/http"
	"testing"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/dynamic/dynamiclister"
	"k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/tools/cache"
	k8smetrics "k8s.io/component-base/metrics"

	volumesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1"

	popv1beta1 "github.com/kubernetes-csi/volume-data-source-validator/client/apis/volumepopulator/v1beta1"
)

type FakeMetricsManager struct{}

func (*FakeMetricsManager) PrepareMetricsPath(mux *http.ServeMux, pattern string, logger promhttp.Logger) error {
	return nil
}
func (*FakeMetricsManager) IncrementCount(result string)         {}
func (*FakeMetricsManager) GetRegistry() k8smetrics.KubeRegistry { return nil }

func makeFakeLister(populators ...*popv1beta1.VolumePopulator) dynamiclister.Lister {
	scheme := runtime.NewScheme()
	popv1beta1.AddToScheme(scheme)
	objects := make([]runtime.Object, len(populators))
	for i := range populators {
		objects[i] = populators[i]
	}
	client := fake.NewSimpleDynamicClient(scheme, objects...)
	factory := dynamicinformer.NewDynamicSharedInformerFactory(client, 0)
	informer := factory.ForResource(PopulatorResource).Informer()
	lister := dynamiclister.New(informer.GetIndexer(), PopulatorResource)
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	cache.WaitForCacheSync(stopCh, informer.HasSynced)
	return lister
}

type brokenVolumeLister struct {
}

func (*brokenVolumeLister) List(labels.Selector) ([]*unstructured.Unstructured, error) {
	return nil, errors.New("failed")
}

func (*brokenVolumeLister) Get(string) (*unstructured.Unstructured, error) {
	return nil, errors.New("failed")
}

func (*brokenVolumeLister) Namespace(string) dynamiclister.NamespaceLister {
	return nil
}

func TestValidateGroupKind(t *testing.T) {
	ctrl := new(populatorController)
	ctrl.metrics = new(FakeMetricsManager)

	populator := popv1beta1.VolumePopulator{
		TypeMeta: metav1.TypeMeta{
			Kind:       "VolumePopulator",
			APIVersion: "populator.storage.k8s.io",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "valid",
		},
		SourceKind: metav1.GroupKind{
			Group: "valid.storage.k8s.io",
			Kind:  "Valid",
		},
	}
	ctrl.popLister = makeFakeLister(&populator)

	testCases := []struct {
		name  string
		gk    metav1.GroupKind
		valid bool
	}{
		{
			name: "Create PVC data source",
			gk: metav1.GroupKind{
				Group: "",
				Kind:  "PersistentVolumeClaim",
			},
			valid: true,
		},
		{
			name: "Create snapshot data source",
			gk: metav1.GroupKind{
				Group: volumesnapshotv1.GroupName,
				Kind:  "VolumeSnapshot",
			},
			valid: true,
		},
		{
			name: "Create valid data source",
			gk: metav1.GroupKind{
				Group: "valid.storage.k8s.io",
				Kind:  "Valid",
			},
			valid: true,
		},
		{
			name: "Create invalid data source",
			gk: metav1.GroupKind{
				Group: "invalid.storage.k8s.io",
				Kind:  "Invalid",
			},
			valid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			valid, err := ctrl.validateGroupKind(tc.gk)
			if valid != tc.valid {
				t.Errorf(`expected "%v" to equal "%v"`, valid, tc.valid)
			}
			if err != nil {
				t.Errorf(`expected nil error, got "%v"`, err)
			}
		})
	}
}

func TestPopListError(t *testing.T) {
	ctrl := new(populatorController)
	ctrl.metrics = new(FakeMetricsManager)
	ctrl.popLister = new(brokenVolumeLister)

	valid, err := ctrl.validateGroupKind(metav1.GroupKind{
		Group: "valid.storage.k8s.io",
		Kind:  "Valid",
	})
	if valid {
		t.Error("expected invalid")
	}
	if nil == err {
		t.Error("expected error")
	}
	if err.Error() != "failed" {
		t.Errorf(`expected "%v" to equal "failed"`, err)
	}
}
