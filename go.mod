module github.com/kubernetes-csi/volume-data-source-validator

go 1.15

require (
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.10 // indirect
	github.com/kubernetes-csi/csi-lib-utils v0.9.1
	github.com/kubernetes-csi/external-snapshotter/client/v4 v4.0.0
	github.com/kubernetes-csi/volume-data-source-validator/client v0.0.0-00010101000000-000000000000
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	golang.org/x/oauth2 v0.0.0-20201208152858-08078c50e5b5 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	k8s.io/api v0.21.0
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
	k8s.io/klog/v2 v2.8.0
)

replace github.com/kubernetes-csi/volume-data-source-validator/client => ./client
