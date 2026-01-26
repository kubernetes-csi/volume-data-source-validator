# Release notes for v1.6.0

# Changelog since v1.5.0

## Changes by Kind

### Other (Cleanup or Flake)

- Bump k8s dependencies to v1.35.0 ([#189](https://github.com/kubernetes-csi/volume-data-source-validator/pull/189), [@dfajmon](https://github.com/dfajmon))

## Dependencies

### Added
- github.com/Masterminds/semver/v3: [v3.4.0](https://github.com/Masterminds/semver/v3/tree/v3.4.0)
- github.com/go-openapi/swag/cmdutils: [v0.25.4](https://github.com/go-openapi/swag/cmdutils/tree/v0.25.4)
- github.com/go-openapi/swag/conv: [v0.25.4](https://github.com/go-openapi/swag/conv/tree/v0.25.4)
- github.com/go-openapi/swag/fileutils: [v0.25.4](https://github.com/go-openapi/swag/fileutils/tree/v0.25.4)
- github.com/go-openapi/swag/jsonname: [v0.25.4](https://github.com/go-openapi/swag/jsonname/tree/v0.25.4)
- github.com/go-openapi/swag/jsonutils/fixtures_test: [v0.25.4](https://github.com/go-openapi/swag/jsonutils/fixtures_test/tree/v0.25.4)
- github.com/go-openapi/swag/jsonutils: [v0.25.4](https://github.com/go-openapi/swag/jsonutils/tree/v0.25.4)
- github.com/go-openapi/swag/loading: [v0.25.4](https://github.com/go-openapi/swag/loading/tree/v0.25.4)
- github.com/go-openapi/swag/mangling: [v0.25.4](https://github.com/go-openapi/swag/mangling/tree/v0.25.4)
- github.com/go-openapi/swag/netutils: [v0.25.4](https://github.com/go-openapi/swag/netutils/tree/v0.25.4)
- github.com/go-openapi/swag/stringutils: [v0.25.4](https://github.com/go-openapi/swag/stringutils/tree/v0.25.4)
- github.com/go-openapi/swag/typeutils: [v0.25.4](https://github.com/go-openapi/swag/typeutils/tree/v0.25.4)
- github.com/go-openapi/swag/yamlutils: [v0.25.4](https://github.com/go-openapi/swag/yamlutils/tree/v0.25.4)
- github.com/go-openapi/testify/enable/yaml/v2: [v2.0.2](https://github.com/go-openapi/testify/enable/yaml/v2/tree/v2.0.2)
- github.com/go-openapi/testify/v2: [v2.0.2](https://github.com/go-openapi/testify/v2/tree/v2.0.2)
- github.com/golang-jwt/jwt/v5: [v5.3.0](https://github.com/golang-jwt/jwt/v5/tree/v5.3.0)
- github.com/prashantv/gostub: [v1.1.0](https://github.com/prashantv/gostub/tree/v1.1.0)
- go.uber.org/automaxprocs: v1.6.0
- golang.org/x/tools/go/expect: v0.1.0-deprecated
- golang.org/x/tools/go/packages/packagestest: v0.1.1-deprecated

### Changed
- github.com/alecthomas/units: [b94a6e3 → 0f3dac3](https://github.com/alecthomas/units/compare/b94a6e3...0f3dac3)
- github.com/emicklei/go-restful/v3: [v3.12.2 → v3.13.0](https://github.com/emicklei/go-restful/v3/compare/v3.12.2...v3.13.0)
- github.com/go-logr/logr: [v1.4.2 → v1.4.3](https://github.com/go-logr/logr/compare/v1.4.2...v1.4.3)
- github.com/go-openapi/jsonpointer: [v0.21.0 → v0.22.4](https://github.com/go-openapi/jsonpointer/compare/v0.21.0...v0.22.4)
- github.com/go-openapi/jsonreference: [v0.21.0 → v0.21.4](https://github.com/go-openapi/jsonreference/compare/v0.21.0...v0.21.4)
- github.com/go-openapi/swag: [v0.23.0 → v0.25.4](https://github.com/go-openapi/swag/compare/v0.23.0...v0.25.4)
- github.com/golang/protobuf: [v1.5.4 → v1.5.2](https://github.com/golang/protobuf/compare/v1.5.4...v1.5.2)
- github.com/google/gnostic-models: [v0.7.0 → v0.7.1](https://github.com/google/gnostic-models/compare/v0.7.0...v0.7.1)
- github.com/google/pprof: [d1b30fe → 27863c8](https://github.com/google/pprof/compare/d1b30fe...27863c8)
- github.com/kubernetes-csi/csi-lib-utils: [v0.20.0 → v0.23.1](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.20.0...v0.23.1)
- github.com/kubernetes-csi/external-snapshotter/client/v6: [v6.2.0 → v6.3.0](https://github.com/kubernetes-csi/external-snapshotter/client/v6/compare/v6.2.0...v6.3.0)
- github.com/onsi/ginkgo/v2: [v2.21.0 → v2.27.2](https://github.com/onsi/ginkgo/v2/compare/v2.21.0...v2.27.2)
- github.com/onsi/gomega: [v1.35.1 → v1.38.2](https://github.com/onsi/gomega/compare/v1.35.1...v1.38.2)
- github.com/prometheus/client_golang: [v1.22.0 → v1.23.2](https://github.com/prometheus/client_golang/compare/v1.22.0...v1.23.2)
- github.com/prometheus/client_model: [v0.6.1 → v0.6.2](https://github.com/prometheus/client_model/compare/v0.6.1...v0.6.2)
- github.com/prometheus/common: [v0.62.0 → v0.67.5](https://github.com/prometheus/common/compare/v0.62.0...v0.67.5)
- github.com/prometheus/procfs: [v0.15.1 → v0.19.2](https://github.com/prometheus/procfs/compare/v0.15.1...v0.19.2)
- github.com/rogpeppe/go-internal: [v1.13.1 → v1.14.1](https://github.com/rogpeppe/go-internal/compare/v1.13.1...v1.14.1)
- github.com/spf13/cobra: [v1.9.1 → v1.10.0](https://github.com/spf13/cobra/compare/v1.9.1...v1.10.0)
- github.com/spf13/pflag: [v1.0.6 → v1.0.10](https://github.com/spf13/pflag/compare/v1.0.6...v1.0.10)
- github.com/stretchr/testify: [v1.10.0 → v1.11.1](https://github.com/stretchr/testify/compare/v1.10.0...v1.11.1)
- go.opentelemetry.io/auto/sdk: v1.1.0 → v1.2.1
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.58.0 → v0.61.0
- go.opentelemetry.io/otel/metric: v1.35.0 → v1.39.0
- go.opentelemetry.io/otel/sdk: v1.34.0 → v1.36.0
- go.opentelemetry.io/otel/trace: v1.35.0 → v1.39.0
- go.opentelemetry.io/otel: v1.35.0 → v1.39.0
- go.yaml.in/yaml/v2: v2.4.2 → v2.4.3
- golang.org/x/crypto: v0.36.0 → v0.47.0
- golang.org/x/mod: v0.20.0 → v0.31.0
- golang.org/x/net: v0.38.0 → v0.49.0
- golang.org/x/oauth2: v0.27.0 → v0.34.0
- golang.org/x/sync: v0.12.0 → v0.19.0
- golang.org/x/sys: v0.31.0 → v0.40.0
- golang.org/x/term: v0.30.0 → v0.39.0
- golang.org/x/text: v0.23.0 → v0.33.0
- golang.org/x/time: v0.9.0 → v0.14.0
- golang.org/x/tools: v0.26.0 → v0.40.0
- google.golang.org/genproto/googleapis/rpc: a0af3ef → 200df99
- google.golang.org/grpc: v1.72.1 → v1.72.2
- google.golang.org/protobuf: v1.36.5 → v1.36.11
- gopkg.in/evanphx/json-patch.v4: v4.12.0 → v4.13.0
- k8s.io/api: v0.34.0 → v0.35.0
- k8s.io/apimachinery: v0.34.1 → v0.35.0
- k8s.io/client-go: v0.34.0 → v0.35.0
- k8s.io/code-generator: v0.25.2 → v0.28.1
- k8s.io/component-base: v0.34.0 → v0.35.0
- k8s.io/gengo: 3913671 → fad74ee
- k8s.io/kube-openapi: f3f2b99 → 4e65d59
- k8s.io/utils: 4c0f3b2 → 914a6e7
- sigs.k8s.io/json: cfa47c3 → 2d32026
- sigs.k8s.io/structured-merge-diff/v4: v4.5.0 → v4.2.3
- sigs.k8s.io/structured-merge-diff/v6: v6.3.0 → v6.3.1

### Removed
- github.com/kisielk/errcheck: [v1.5.0](https://github.com/kisielk/errcheck/tree/v1.5.0)
- github.com/kisielk/gotool: [v1.0.0](https://github.com/kisielk/gotool/tree/v1.0.0)
- github.com/yuin/goldmark: [v1.2.1](https://github.com/yuin/goldmark/tree/v1.2.1)
- golang.org/x/xerrors: 5ec99f8
