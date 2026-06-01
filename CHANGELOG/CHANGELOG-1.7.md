# Release notes for v1.7.0

# Changelog since v1.6.0

## Changes by Kind

### Other (Cleanup or Flake)

- Bump k8s dependencies to v1.36.1 ([#217](https://github.com/kubernetes-csi/volume-data-source-validator/pull/217), [@dfajmon](https://github.com/dfajmon))

## Dependencies

### Added
- github.com/cenkalti/backoff/v5: [v5.0.3](https://github.com/cenkalti/backoff/v5/tree/v5.0.3)
- k8s.io/streaming: v0.36.1

### Changed
- github.com/grpc-ecosystem/grpc-gateway/v2: [v2.26.3 → v2.27.7](https://github.com/grpc-ecosystem/grpc-gateway/v2/compare/v2.26.3...v2.27.7)
- github.com/kubernetes-csi/csi-lib-utils: [v0.23.1 → v0.24.0](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.23.1...v0.24.0)
- github.com/moby/spdystream: [v0.5.0 → v0.5.1](https://github.com/moby/spdystream/compare/v0.5.0...v0.5.1)
- github.com/spf13/cobra: [v1.10.0 → v1.10.2](https://github.com/spf13/cobra/compare/v1.10.0...v1.10.2)
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.61.0 → v0.65.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc: v1.34.0 → v1.40.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace: v1.34.0 → v1.40.0
- go.opentelemetry.io/otel/metric: v1.39.0 → v1.41.0
- go.opentelemetry.io/otel/sdk: v1.36.0 → v1.40.0
- go.opentelemetry.io/otel/trace: v1.39.0 → v1.41.0
- go.opentelemetry.io/otel: v1.39.0 → v1.41.0
- go.opentelemetry.io/proto/otlp: v1.5.0 → v1.9.0
- go.uber.org/zap: v1.27.0 → v1.27.1
- google.golang.org/genproto/googleapis/api: a0af3ef → 8636f87
- google.golang.org/genproto/googleapis/rpc: 200df99 → 8636f87
- google.golang.org/grpc: v1.72.2 → v1.79.3
- google.golang.org/protobuf: v1.36.11 → f2248ac
- k8s.io/api: v0.35.0 → v0.36.1
- k8s.io/apimachinery: v0.35.0 → v0.36.1
- k8s.io/client-go: v0.35.0 → v0.36.1
- k8s.io/component-base: v0.35.0 → v0.36.1
- k8s.io/klog/v2: v2.130.1 → v2.140.0
- k8s.io/kube-openapi: 4e65d59 → 43fb72c
- k8s.io/utils: 914a6e7 → b8788ab
- sigs.k8s.io/structured-merge-diff/v6: v6.3.1 → v6.3.2

### Removed
- github.com/Masterminds/semver/v3: [v3.4.0](https://github.com/Masterminds/semver/v3/tree/v3.4.0)
- github.com/armon/go-socks5: [e753329](https://github.com/armon/go-socks5/tree/e753329)
- github.com/cenkalti/backoff/v4: [v4.3.0](https://github.com/cenkalti/backoff/v4/tree/v4.3.0)
- github.com/go-task/slim-sprig/v3: [v3.0.0](https://github.com/go-task/slim-sprig/v3/tree/v3.0.0)
- github.com/google/pprof: [27863c8](https://github.com/google/pprof/tree/27863c8)
- github.com/gregjones/httpcache: [901d907](https://github.com/gregjones/httpcache/tree/901d907)
- github.com/onsi/ginkgo/v2: [v2.27.2](https://github.com/onsi/ginkgo/v2/tree/v2.27.2)
- github.com/onsi/gomega: [v1.38.2](https://github.com/onsi/gomega/tree/v1.38.2)
