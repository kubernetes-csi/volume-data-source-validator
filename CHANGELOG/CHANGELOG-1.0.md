# Release notes for v1.0.0

## Changes by Kind

### API Change

- Renamed v1alpha1 API to v1beta1. The API is the same, but users must use the new version. ([#17](https://github.com/kubernetes-csi/volume-data-source-validator/pull/17), [@bswartz](https://github.com/bswartz))

### Feature

- The controller now exposes metrics covering the counts and results of validation operations. ([#14](https://github.com/kubernetes-csi/volume-data-source-validator/pull/14), [@bswartz](https://github.com/bswartz))

### Bug or Regression

- By default, metrics are now exposed on port 8080. ([#26](https://github.com/kubernetes-csi/volume-data-source-validator/pull/26), [@bswartz](https://github.com/bswartz))
- Fixed a bug that caused volumes to get re-validated periodically, generating useless log messages. ([#25](https://github.com/kubernetes-csi/volume-data-source-validator/pull/25), [@bswartz](https://github.com/bswartz))
- The yaml for the CRD now correctly installs VolumePopulator as a cluster-scoped object. ([#24](https://github.com/kubernetes-csi/volume-data-source-validator/pull/24), [@bswartz](https://github.com/bswartz))

### Uncategorized

- Kuberenetes client dependencies are updated to v1.23 with this release ([#21](https://github.com/kubernetes-csi/volume-data-source-validator/pull/21), [@humblec](https://github.com/humblec))

## Dependencies

### Added
- cloud.google.com/go/firestore: v1.1.0
- github.com/OneOfOne/xxhash: [v1.2.2](https://github.com/OneOfOne/xxhash/tree/v1.2.2)
- github.com/antihax/optional: [v1.0.0](https://github.com/antihax/optional/tree/v1.0.0)
- github.com/armon/circbuf: [bbbad09](https://github.com/armon/circbuf/tree/bbbad09)
- github.com/armon/go-metrics: [f0300d1](https://github.com/armon/go-metrics/tree/f0300d1)
- github.com/armon/go-radix: [7fddfc3](https://github.com/armon/go-radix/tree/7fddfc3)
- github.com/benbjohnson/clock: [v1.1.0](https://github.com/benbjohnson/clock/tree/v1.1.0)
- github.com/bgentry/speakeasy: [v0.1.0](https://github.com/bgentry/speakeasy/tree/v0.1.0)
- github.com/bketelsen/crypt: [v0.0.4](https://github.com/bketelsen/crypt/tree/v0.0.4)
- github.com/cespare/xxhash: [v1.1.0](https://github.com/cespare/xxhash/tree/v1.1.0)
- github.com/cncf/xds/go: [fbca930](https://github.com/cncf/xds/go/tree/fbca930)
- github.com/coreos/go-semver: [v0.3.0](https://github.com/coreos/go-semver/tree/v0.3.0)
- github.com/coreos/go-systemd/v22: [v22.3.2](https://github.com/coreos/go-systemd/v22/tree/v22.3.2)
- github.com/cpuguy83/go-md2man/v2: [v2.0.0](https://github.com/cpuguy83/go-md2man/v2/tree/v2.0.0)
- github.com/fatih/color: [v1.7.0](https://github.com/fatih/color/tree/v1.7.0)
- github.com/felixge/httpsnoop: [v1.0.1](https://github.com/felixge/httpsnoop/tree/v1.0.1)
- github.com/getkin/kin-openapi: [v0.76.0](https://github.com/getkin/kin-openapi/tree/v0.76.0)
- github.com/go-kit/log: [v0.1.0](https://github.com/go-kit/log/tree/v0.1.0)
- github.com/go-logr/zapr: [v1.2.0](https://github.com/go-logr/zapr/tree/v1.2.0)
- github.com/godbus/dbus/v5: [v5.0.4](https://github.com/godbus/dbus/v5/tree/v5.0.4)
- github.com/gopherjs/gopherjs: [0766667](https://github.com/gopherjs/gopherjs/tree/0766667)
- github.com/gorilla/mux: [v1.8.0](https://github.com/gorilla/mux/tree/v1.8.0)
- github.com/grpc-ecosystem/grpc-gateway: [v1.16.0](https://github.com/grpc-ecosystem/grpc-gateway/tree/v1.16.0)
- github.com/hashicorp/consul/api: [v1.1.0](https://github.com/hashicorp/consul/api/tree/v1.1.0)
- github.com/hashicorp/consul/sdk: [v0.1.1](https://github.com/hashicorp/consul/sdk/tree/v0.1.1)
- github.com/hashicorp/errwrap: [v1.0.0](https://github.com/hashicorp/errwrap/tree/v1.0.0)
- github.com/hashicorp/go-cleanhttp: [v0.5.1](https://github.com/hashicorp/go-cleanhttp/tree/v0.5.1)
- github.com/hashicorp/go-immutable-radix: [v1.0.0](https://github.com/hashicorp/go-immutable-radix/tree/v1.0.0)
- github.com/hashicorp/go-msgpack: [v0.5.3](https://github.com/hashicorp/go-msgpack/tree/v0.5.3)
- github.com/hashicorp/go-multierror: [v1.0.0](https://github.com/hashicorp/go-multierror/tree/v1.0.0)
- github.com/hashicorp/go-rootcerts: [v1.0.0](https://github.com/hashicorp/go-rootcerts/tree/v1.0.0)
- github.com/hashicorp/go-sockaddr: [v1.0.0](https://github.com/hashicorp/go-sockaddr/tree/v1.0.0)
- github.com/hashicorp/go-syslog: [v1.0.0](https://github.com/hashicorp/go-syslog/tree/v1.0.0)
- github.com/hashicorp/go-uuid: [v1.0.1](https://github.com/hashicorp/go-uuid/tree/v1.0.1)
- github.com/hashicorp/go.net: [v0.0.1](https://github.com/hashicorp/go.net/tree/v0.0.1)
- github.com/hashicorp/hcl: [v1.0.0](https://github.com/hashicorp/hcl/tree/v1.0.0)
- github.com/hashicorp/logutils: [v1.0.0](https://github.com/hashicorp/logutils/tree/v1.0.0)
- github.com/hashicorp/mdns: [v1.0.0](https://github.com/hashicorp/mdns/tree/v1.0.0)
- github.com/hashicorp/memberlist: [v0.1.3](https://github.com/hashicorp/memberlist/tree/v0.1.3)
- github.com/hashicorp/serf: [v0.8.2](https://github.com/hashicorp/serf/tree/v0.8.2)
- github.com/inconshreveable/mousetrap: [v1.0.0](https://github.com/inconshreveable/mousetrap/tree/v1.0.0)
- github.com/jpillora/backoff: [v1.0.0](https://github.com/jpillora/backoff/tree/v1.0.0)
- github.com/jtolds/gls: [v4.20.0+incompatible](https://github.com/jtolds/gls/tree/v4.20.0)
- github.com/kr/fs: [v0.1.0](https://github.com/kr/fs/tree/v0.1.0)
- github.com/magiconair/properties: [v1.8.5](https://github.com/magiconair/properties/tree/v1.8.5)
- github.com/mattn/go-colorable: [v0.0.9](https://github.com/mattn/go-colorable/tree/v0.0.9)
- github.com/mattn/go-isatty: [v0.0.3](https://github.com/mattn/go-isatty/tree/v0.0.3)
- github.com/miekg/dns: [v1.0.14](https://github.com/miekg/dns/tree/v1.0.14)
- github.com/mitchellh/cli: [v1.0.0](https://github.com/mitchellh/cli/tree/v1.0.0)
- github.com/mitchellh/go-homedir: [v1.0.0](https://github.com/mitchellh/go-homedir/tree/v1.0.0)
- github.com/mitchellh/go-testing-interface: [v1.0.0](https://github.com/mitchellh/go-testing-interface/tree/v1.0.0)
- github.com/mitchellh/gox: [v0.4.0](https://github.com/mitchellh/gox/tree/v0.4.0)
- github.com/mitchellh/iochan: [v1.0.0](https://github.com/mitchellh/iochan/tree/v1.0.0)
- github.com/pascaldekloe/goe: [57f6aae](https://github.com/pascaldekloe/goe/tree/57f6aae)
- github.com/pelletier/go-toml: [v1.9.3](https://github.com/pelletier/go-toml/tree/v1.9.3)
- github.com/pkg/sftp: [v1.10.1](https://github.com/pkg/sftp/tree/v1.10.1)
- github.com/posener/complete: [v1.1.1](https://github.com/posener/complete/tree/v1.1.1)
- github.com/rogpeppe/fastuuid: [v1.2.0](https://github.com/rogpeppe/fastuuid/tree/v1.2.0)
- github.com/russross/blackfriday/v2: [v2.0.1](https://github.com/russross/blackfriday/v2/tree/v2.0.1)
- github.com/ryanuber/columnize: [9b3edd6](https://github.com/ryanuber/columnize/tree/9b3edd6)
- github.com/sean-/seed: [e2103e2](https://github.com/sean-/seed/tree/e2103e2)
- github.com/shurcooL/sanitized_anchor_name: [v1.0.0](https://github.com/shurcooL/sanitized_anchor_name/tree/v1.0.0)
- github.com/smartystreets/assertions: [b2de0cb](https://github.com/smartystreets/assertions/tree/b2de0cb)
- github.com/smartystreets/goconvey: [v1.6.4](https://github.com/smartystreets/goconvey/tree/v1.6.4)
- github.com/spaolacci/murmur3: [f09979e](https://github.com/spaolacci/murmur3/tree/f09979e)
- github.com/spf13/cast: [v1.3.1](https://github.com/spf13/cast/tree/v1.3.1)
- github.com/spf13/cobra: [v1.2.1](https://github.com/spf13/cobra/tree/v1.2.1)
- github.com/spf13/jwalterweatherman: [v1.1.0](https://github.com/spf13/jwalterweatherman/tree/v1.1.0)
- github.com/spf13/viper: [v1.8.1](https://github.com/spf13/viper/tree/v1.8.1)
- github.com/subosito/gotenv: [v1.2.0](https://github.com/subosito/gotenv/tree/v1.2.0)
- go.etcd.io/etcd/api/v3: v3.5.0
- go.etcd.io/etcd/client/pkg/v3: v3.5.0
- go.etcd.io/etcd/client/v2: v2.305.0
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.20.0
- go.opentelemetry.io/contrib: v0.20.0
- go.opentelemetry.io/otel/exporters/otlp: v0.20.0
- go.opentelemetry.io/otel/metric: v0.20.0
- go.opentelemetry.io/otel/oteltest: v0.20.0
- go.opentelemetry.io/otel/sdk/export/metric: v0.20.0
- go.opentelemetry.io/otel/sdk/metric: v0.20.0
- go.opentelemetry.io/otel/sdk: v0.20.0
- go.opentelemetry.io/otel/trace: v0.20.0
- go.opentelemetry.io/otel: v0.20.0
- go.opentelemetry.io/proto/otlp: v0.7.0
- go.uber.org/goleak: v1.1.10
- gopkg.in/ini.v1: v1.62.0
- sigs.k8s.io/json: c049b76

### Changed
- cloud.google.com/go: v0.65.0 → v0.81.0
- github.com/Azure/go-ansiterm: [d6e3b33 → d185dfc](https://github.com/Azure/go-ansiterm/compare/d6e3b33...d185dfc)
- github.com/alecthomas/units: [c3de453 → f65c72e](https://github.com/alecthomas/units/compare/c3de453...f65c72e)
- github.com/blang/semver: [v3.5.0+incompatible → v3.5.1+incompatible](https://github.com/blang/semver/compare/v3.5.0...v3.5.1)
- github.com/cncf/udpa/go: [269d4d4 → 5459f2c](https://github.com/cncf/udpa/go/compare/269d4d4...5459f2c)
- github.com/creack/pty: [v1.1.9 → v1.1.11](https://github.com/creack/pty/compare/v1.1.9...v1.1.11)
- github.com/envoyproxy/go-control-plane: [v0.9.4 → 63b5d3c](https://github.com/envoyproxy/go-control-plane/compare/v0.9.4...63b5d3c)
- github.com/evanphx/json-patch: [v4.11.0+incompatible → v4.12.0+incompatible](https://github.com/evanphx/json-patch/compare/v4.11.0...v4.12.0)
- github.com/ghodss/yaml: [73d445a → v1.0.0](https://github.com/ghodss/yaml/compare/73d445a...v1.0.0)
- github.com/go-logfmt/logfmt: [v0.4.0 → v0.5.0](https://github.com/go-logfmt/logfmt/compare/v0.4.0...v0.5.0)
- github.com/go-logr/logr: [v0.4.0 → v1.2.0](https://github.com/go-logr/logr/compare/v0.4.0...v1.2.0)
- github.com/go-openapi/jsonpointer: [v0.19.3 → v0.19.5](https://github.com/go-openapi/jsonpointer/compare/v0.19.3...v0.19.5)
- github.com/golang/mock: [v1.4.4 → v1.5.0](https://github.com/golang/mock/compare/v1.4.4...v1.5.0)
- github.com/google/martian/v3: [v3.0.0 → v3.1.0](https://github.com/google/martian/v3/compare/v3.0.0...v3.1.0)
- github.com/google/pprof: [1a94d86 → cbba55b](https://github.com/google/pprof/compare/1a94d86...cbba55b)
- github.com/ianlancetaylor/demangle: [5e5cf60 → 28f6c0f](https://github.com/ianlancetaylor/demangle/compare/5e5cf60...28f6c0f)
- github.com/json-iterator/go: [v1.1.11 → v1.1.12](https://github.com/json-iterator/go/compare/v1.1.11...v1.1.12)
- github.com/julienschmidt/httprouter: [v1.2.0 → v1.3.0](https://github.com/julienschmidt/httprouter/compare/v1.2.0...v1.3.0)
- github.com/mitchellh/mapstructure: [v1.1.2 → v1.4.1](https://github.com/mitchellh/mapstructure/compare/v1.1.2...v1.4.1)
- github.com/moby/term: [672ec06 → 9d4ed18](https://github.com/moby/term/compare/672ec06...9d4ed18)
- github.com/modern-go/reflect2: [v1.0.1 → v1.0.2](https://github.com/modern-go/reflect2/compare/v1.0.1...v1.0.2)
- github.com/mwitkow/go-conntrack: [cc309e4 → 2f06839](https://github.com/mwitkow/go-conntrack/compare/cc309e4...2f06839)
- github.com/prometheus/client_golang: [v1.7.1 → v1.11.0](https://github.com/prometheus/client_golang/compare/v1.7.1...v1.11.0)
- github.com/prometheus/common: [v0.10.0 → v0.28.0](https://github.com/prometheus/common/compare/v0.10.0...v0.28.0)
- github.com/prometheus/procfs: [v0.1.3 → v0.6.0](https://github.com/prometheus/procfs/compare/v0.1.3...v0.6.0)
- github.com/spf13/afero: [v1.2.2 → v1.6.0](https://github.com/spf13/afero/compare/v1.2.2...v1.6.0)
- github.com/yuin/goldmark: [v1.2.1 → v1.4.0](https://github.com/yuin/goldmark/compare/v1.2.1...v1.4.0)
- go.opencensus.io: v0.22.4 → v0.23.0
- go.uber.org/atomic: v1.4.0 → v1.7.0
- go.uber.org/multierr: v1.1.0 → v1.6.0
- go.uber.org/zap: v1.10.0 → v1.19.0
- golang.org/x/crypto: 5ea612d → 32db794
- golang.org/x/lint: 738671d → 6edffad
- golang.org/x/mod: v0.3.0 → v0.4.2
- golang.org/x/net: 37e1c6a → 491a49a
- golang.org/x/oauth2: 08078c5 → 2bc19b1
- golang.org/x/sync: 67f06af → 036812b
- golang.org/x/sys: 59db8d7 → f4d4317
- golang.org/x/term: 6a3ed07 → 6886f2d
- golang.org/x/text: v0.3.6 → v0.3.7
- golang.org/x/tools: 113979e → d4cc65f
- google.golang.org/api: v0.30.0 → v0.44.0
- google.golang.org/genproto: 1ed22bb → fe13028
- google.golang.org/grpc: v1.31.0 → v1.40.0
- google.golang.org/protobuf: v1.26.0 → v1.27.1
- gotest.tools/v3: v3.0.2 → v3.0.3
- k8s.io/api: v0.22.0 → v0.23.4
- k8s.io/apimachinery: v0.22.0 → v0.23.4
- k8s.io/client-go: v0.22.0 → v0.23.4
- k8s.io/component-base: v0.19.0 → v0.23.4
- k8s.io/gengo: 8167cfd → 485abfe
- k8s.io/klog/v2: v2.9.0 → v2.30.0
- k8s.io/kube-openapi: 9528897 → e816edb
- k8s.io/utils: 4b05e18 → 6203023
- sigs.k8s.io/structured-merge-diff/v4: v4.1.2 → v4.2.1

### Removed
_Nothing has changed._
