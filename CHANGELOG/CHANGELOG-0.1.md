# Release notes for v0.1.0

## Changes by Kind

### Feature

- Implement validation of data sources for PVC based on the VolumePopulators framework.

### Other (Cleanup or Flake)

- Image name updated to volume-data-source-manager. ([#11](https://github.com/kubernetes-csi/volume-data-source-validator/pull/11), [@bswartz](https://github.com/bswartz))

### Uncategorized

- Kubernetes v1.22 or later is required, and the AnyVolumeDataSource feature gate must be enabled ([#5](https://github.com/kubernetes-csi/volume-data-source-validator/pull/5), [@bswartz](https://github.com/bswartz))
- Removed generated clients. Please use the dynamic client when writing controllers to interact with VolumePopulator objects. ([#2](https://github.com/kubernetes-csi/volume-data-source-validator/pull/2), [@bswartz](https://github.com/bswartz))

## Dependencies

### Added
- github.com/Azure/go-autorest: [v14.2.0+incompatible](https://github.com/Azure/go-autorest/tree/v14.2.0)
- github.com/asaskevich/govalidator: [f61b66f](https://github.com/asaskevich/govalidator/tree/f61b66f)
- github.com/creack/pty: [v1.1.9](https://github.com/creack/pty/tree/v1.1.9)
- github.com/form3tech-oss/jwt-go: [v3.2.3+incompatible](https://github.com/form3tech-oss/jwt-go/tree/v3.2.3)
- github.com/gorilla/websocket: [v1.4.2](https://github.com/gorilla/websocket/tree/v1.4.2)
- github.com/kubernetes-csi/external-snapshotter/client/v4: [v4.0.0](https://github.com/kubernetes-csi/external-snapshotter/client/v4/tree/v4.0.0)
- github.com/mitchellh/mapstructure: [v1.1.2](https://github.com/mitchellh/mapstructure/tree/v1.1.2)
- github.com/moby/spdystream: [v0.2.0](https://github.com/moby/spdystream/tree/v0.2.0)
- github.com/niemeyer/pretty: [a10e7ca](https://github.com/niemeyer/pretty/tree/a10e7ca)

### Changed
- github.com/Azure/go-autorest/autorest/adal: [v0.8.2 → v0.9.13](https://github.com/Azure/go-autorest/autorest/adal/compare/v0.8.2...v0.9.13)
- github.com/Azure/go-autorest/autorest/date: [v0.2.0 → v0.3.0](https://github.com/Azure/go-autorest/autorest/date/compare/v0.2.0...v0.3.0)
- github.com/Azure/go-autorest/autorest/mocks: [v0.3.0 → v0.4.1](https://github.com/Azure/go-autorest/autorest/mocks/compare/v0.3.0...v0.4.1)
- github.com/Azure/go-autorest/autorest: [v0.9.6 → v0.11.18](https://github.com/Azure/go-autorest/autorest/compare/v0.9.6...v0.11.18)
- github.com/Azure/go-autorest/logger: [v0.1.0 → v0.2.1](https://github.com/Azure/go-autorest/logger/compare/v0.1.0...v0.2.1)
- github.com/Azure/go-autorest/tracing: [v0.5.0 → v0.6.0](https://github.com/Azure/go-autorest/tracing/compare/v0.5.0...v0.6.0)
- github.com/evanphx/json-patch: [v4.9.0+incompatible → v4.11.0+incompatible](https://github.com/evanphx/json-patch/compare/v4.9.0...v4.11.0)
- github.com/gogo/protobuf: [v1.3.1 → v1.3.2](https://github.com/gogo/protobuf/compare/v1.3.1...v1.3.2)
- github.com/golang/groupcache: [8c9f03a → 41bb18b](https://github.com/golang/groupcache/compare/8c9f03a...41bb18b)
- github.com/golang/protobuf: [v1.4.3 → v1.5.2](https://github.com/golang/protobuf/compare/v1.4.3...v1.5.2)
- github.com/google/btree: [v1.0.0 → v1.0.1](https://github.com/google/btree/compare/v1.0.0...v1.0.1)
- github.com/google/go-cmp: [v0.5.2 → v0.5.5](https://github.com/google/go-cmp/compare/v0.5.2...v0.5.5)
- github.com/googleapis/gnostic: [v0.5.1 → v0.5.5](https://github.com/googleapis/gnostic/compare/v0.5.1...v0.5.5)
- github.com/hashicorp/golang-lru: [v0.5.4 → v0.5.1](https://github.com/hashicorp/golang-lru/compare/v0.5.4...v0.5.1)
- github.com/json-iterator/go: [v1.1.10 → v1.1.11](https://github.com/json-iterator/go/compare/v1.1.10...v1.1.11)
- github.com/kisielk/errcheck: [v1.2.0 → v1.5.0](https://github.com/kisielk/errcheck/compare/v1.2.0...v1.5.0)
- github.com/kr/text: [v0.1.0 → v0.2.0](https://github.com/kr/text/compare/v0.1.0...v0.2.0)
- github.com/stretchr/testify: [v1.5.1 → v1.7.0](https://github.com/stretchr/testify/compare/v1.5.1...v1.7.0)
- github.com/yuin/goldmark: [v1.1.32 → v1.2.1](https://github.com/yuin/goldmark/compare/v1.1.32...v1.2.1)
- golang.org/x/crypto: 5f87f34 → 5ea612d
- golang.org/x/net: ac852fb → 37e1c6a
- golang.org/x/sync: 6e8e738 → 67f06af
- golang.org/x/sys: aec9a39 → 59db8d7
- golang.org/x/term: 2321bbc → 6a3ed07
- golang.org/x/text: v0.3.3 → v0.3.6
- golang.org/x/time: 7e3f01d → 1f47c86
- golang.org/x/tools: b303f43 → 113979e
- google.golang.org/genproto: 8632dd7 → 1ed22bb
- google.golang.org/protobuf: v1.25.0 → v1.26.0
- gopkg.in/check.v1: 41f04d3 → 8fa4692
- gopkg.in/yaml.v3: eeeca48 → 496545a
- k8s.io/api: v0.19.9 → v0.22.0
- k8s.io/apimachinery: v0.19.9 → v0.22.0
- k8s.io/client-go: v0.19.9 → v0.22.0
- k8s.io/code-generator: v0.19.9 → v0.19.0
- k8s.io/klog/v2: v2.8.0 → v2.9.0
- k8s.io/kube-openapi: 6aeccd4 → 9528897
- k8s.io/utils: 4140de9 → 4b05e18
- sigs.k8s.io/structured-merge-diff/v4: v4.0.1 → v4.1.2

### Removed
- github.com/kubernetes-csi/external-snapshotter/client/v3: [v3.0.0](https://github.com/kubernetes-csi/external-snapshotter/client/v3/tree/v3.0.0)
