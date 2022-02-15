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

package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	k8smetrics "k8s.io/component-base/metrics"
)

const (
	subSystem   = "volume_data_source_validator"
	labelResult = "result"

	DataSourceEmptyResultName     = "empty"
	DataSourcePVCResultName       = "pvc"
	DataSourceSnapshotResultName  = "snapshot"
	DataSourcePopulatorResultName = "populator"
	DataSourceInvalidResultName   = "invalid"
	DataSourceErrorResultName     = "error"
)

type MetricsManager interface {
	// PrepareMetricsPath prepares the metrics path the specified pattern for
	// metrics managed by this MetricsManager.
	// If the "pattern" is empty (i.e., ""), it will not be registered.
	// An error will be returned if there is any.
	PrepareMetricsPath(mux *http.ServeMux, pattern string, logger promhttp.Logger) error

	// IncrementCount records a metric point for a validation operation.
	// result - the result of the validation operation.
	IncrementCount(result string)

	// GetRegistry() returns the metrics.KubeRegistry used by this metrics manager.
	GetRegistry() k8smetrics.KubeRegistry
}

type operationMetricsManager struct {
	// registry is a wrapper around Prometheus Registry
	registry k8smetrics.KubeRegistry

	// opResultMetrics is a COunter metrics for operation results
	opResultMetrics *k8smetrics.CounterVec
}

// NewMetricsManager creates a new MetricsManager instance
func NewMetricsManager() MetricsManager {
	mgr := new(operationMetricsManager)
	mgr.init()
	return mgr
}

// RecordMetrics emits operation metrics
func (opMgr *operationMetricsManager) IncrementCount(result string) {
	opMgr.opResultMetrics.WithLabelValues(result).Inc()
}

func (opMgr *operationMetricsManager) init() {
	opMgr.registry = k8smetrics.NewKubeRegistry()
	k8smetrics.RegisterProcessStartTime(opMgr.registry.Register)
	opMgr.opResultMetrics = k8smetrics.NewCounterVec(
		&k8smetrics.CounterOpts{
			Subsystem: subSystem,
			Name:      "operation_count",
			Help:      "Number of validations operations by result",
		},
		[]string{labelResult},
	)
	opMgr.registry.MustRegister(opMgr.opResultMetrics)
}

func (opMgr *operationMetricsManager) PrepareMetricsPath(mux *http.ServeMux, pattern string, logger promhttp.Logger) error {
	mux.Handle(pattern, k8smetrics.HandlerFor(
		opMgr.registry,
		k8smetrics.HandlerOpts{
			ErrorLog:      logger,
			ErrorHandling: k8smetrics.ContinueOnError,
		}))

	return nil
}

func (opMgr *operationMetricsManager) GetRegistry() k8smetrics.KubeRegistry {
	return opMgr.registry
}
