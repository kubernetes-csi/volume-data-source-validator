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

package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	coreinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	"github.com/kubernetes-csi/csi-lib-utils/leaderelection"

	popv1beta1 "github.com/kubernetes-csi/volume-data-source-validator/client/apis/volumepopulator/v1beta1"

	popcontroller "github.com/kubernetes-csi/volume-data-source-validator/pkg/data-source-validator"
	"github.com/kubernetes-csi/volume-data-source-validator/pkg/metrics"
)

// Command line flags
var (
	kubeconfig   = flag.String("kubeconfig", "", "Absolute path to the kubeconfig file. Required only when running out of cluster.")
	resyncPeriod = flag.Duration("resync-period", 60*time.Second, "Resync interval of the controller.")
	showVersion  = flag.Bool("version", false, "Show version.")
	threads      = flag.Int("worker-threads", 10, "Number of worker threads.")

	leaderElection              = flag.Bool("leader-election", false, "Enables leader election.")
	leaderElectionNamespace     = flag.String("leader-election-namespace", "", "The namespace where the leader election resource exists. Defaults to the pod namespace if not set.")
	leaderElectionLeaseDuration = flag.Duration("leader-election-lease-duration", 15*time.Second, "Duration, in seconds, that non-leader candidates will wait to force acquire leadership. Defaults to 15 seconds.")
	leaderElectionRenewDeadline = flag.Duration("leader-election-renew-deadline", 10*time.Second, "Duration, in seconds, that the acting leader will retry refreshing leadership before giving up. Defaults to 10 seconds.")
	leaderElectionRetryPeriod   = flag.Duration("leader-election-retry-period", 5*time.Second, "Duration, in seconds, the LeaderElector clients should wait between tries of actions. Defaults to 5 seconds.")

	httpEndpoint = flag.String("http-endpoint", "", "The TCP network address where the HTTP server for diagnostics, including metrics and leader election health check, will listen (example: `:8080`). The default is empty string, which means the server is disabled.")
	metricsPath  = flag.String("metrics-path", "/metrics", "The HTTP path where prometheus metrics will be exposed. Default is `/metrics`.")
)

var (
	version = "unknown"
)

func main() {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "true")
	flag.Parse()

	if *showVersion {
		fmt.Println(os.Args[0], version)
		os.Exit(0)
	}
	klog.Infof("Version: %s", version)

	// Create the client config. Use kubeconfig if given, otherwise assume in-cluster.
	config, err := buildConfig(*kubeconfig)
	if err != nil {
		klog.Error(err.Error())
		os.Exit(1)
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Error(err.Error())
		os.Exit(1)
	}

	dynClient, err := dynamic.NewForConfig(config)
	if nil != err {
		klog.Fatalf("Failed to create dynamic client: %v", err)
	}

	coreFactory := coreinformers.NewSharedInformerFactory(kubeClient, *resyncPeriod)
	dynFactory := dynamicinformer.NewDynamicSharedInformerFactory(dynClient, *resyncPeriod)

	// Create and register metrics manager
	metricsManager := metrics.NewMetricsManager()
	wg := &sync.WaitGroup{}

	mux := http.NewServeMux()
	if *httpEndpoint != "" {
		err := metricsManager.PrepareMetricsPath(mux, *metricsPath, promklog{})
		if err != nil {
			klog.Errorf("Failed to prepare metrics path: %s", err.Error())
			os.Exit(1)
		}
		klog.Infof("Metrics path successfully registered at %s", *metricsPath)
	}
	popv1beta1.AddToScheme(scheme.Scheme)

	klog.V(2).Infof("Start NewDataSourceValidator with kubeconfig [%s] resyncPeriod [%+v]", *kubeconfig, *resyncPeriod)

	ctrl := popcontroller.NewDataSourceValidator(
		dynClient,
		kubeClient,
		dynFactory.ForResource(popcontroller.PopulatorResource).Informer(),
		coreFactory.Core().V1().PersistentVolumeClaims(),
		metricsManager,
		*resyncPeriod,
	)

	run := func(context.Context) {
		// run...
		stopCh := make(chan struct{})
		dynFactory.Start(stopCh)
		coreFactory.Start(stopCh)
		go ctrl.Run(*threads, stopCh)

		// ...until SIGINT
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		close(stopCh)
	}

	// start listening & serving http endpoint if set
	if *httpEndpoint != "" {
		l, err := net.Listen("tcp", *httpEndpoint)
		if err != nil {
			klog.Fatalf("failed to listen on address[%s], error[%v]", *httpEndpoint, err)
		}
		srv := &http.Server{Addr: l.Addr().String(), Handler: mux}
		go func() {
			defer wg.Done()
			if err := srv.Serve(l); err != http.ErrServerClosed {
				klog.Fatalf("failed to start endpoint at:%s/%s, error: %v", *httpEndpoint, *metricsPath, err)
			}
		}()
		klog.Infof("Metrics http server successfully started on %s, %s", *httpEndpoint, *metricsPath)

		defer func() {
			err := srv.Shutdown(context.Background())
			if err != nil {
				klog.Errorf("Failed to shutdown metrics server: %s", err.Error())
			}

			klog.Infof("Metrics server successfully shutdown")
			wg.Done()
		}()
	}

	if !*leaderElection {
		run(context.TODO())
	} else {
		lockName := "data-source-validator-leader"
		// Create a new clientset for leader election to prevent throttling
		// due to populator controller
		leClientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			klog.Fatalf("failed to create leaderelection client: %v", err)
		}
		le := leaderelection.NewLeaderElection(leClientset, lockName, run)
		if *httpEndpoint != "" {
			le.PrepareHealthCheck(mux, leaderelection.DefaultHealthCheckTimeout)
		}

		if *leaderElectionNamespace != "" {
			le.WithNamespace(*leaderElectionNamespace)
		}
		le.WithLeaseDuration(*leaderElectionLeaseDuration)
		le.WithRenewDeadline(*leaderElectionRenewDeadline)
		le.WithRetryPeriod(*leaderElectionRetryPeriod)
		if err := le.Run(); err != nil {
			klog.Fatalf("failed to initialize leader election: %v", err)
		}
	}
}

func buildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}

type promklog struct{}

func (pl promklog) Println(v ...interface{}) {
	klog.Error(v...)
}
