module github.com/kubernetes-csi/volume-data-source-validator

go 1.16

require (
	github.com/imdario/mergo v0.3.10 // indirect
	github.com/kubernetes-csi/csi-lib-utils v0.11.0
	github.com/kubernetes-csi/external-snapshotter/client/v6 v6.0.1
	github.com/kubernetes-csi/volume-data-source-validator/client v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.12.1
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.32.1
	k8s.io/api v0.24.0
	k8s.io/apimachinery v0.24.0
	k8s.io/client-go v0.24.0
	k8s.io/component-base v0.24.0
	k8s.io/klog/v2 v2.60.1
)

replace github.com/kubernetes-csi/volume-data-source-validator/client => ./client
