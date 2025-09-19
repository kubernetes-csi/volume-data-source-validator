# Release notes for v1.5.0

# Changelog since v1.4.0

## Changes by Kind

### Feature

- Update Dependencies for Kubernetes 1.34 ([#184](https://github.com/kubernetes-csi/volume-data-source-validator/pull/184), [@sunnylovestiramisu](https://github.com/sunnylovestiramisu))

### Other (Cleanup or Flake)

- Update kubernetes dependencies to v1.34.0 ([#182](https://github.com/kubernetes-csi/volume-data-source-validator/pull/182), [@dobsonj](https://github.com/dobsonj))

## Dependencies

### Added
- go.yaml.in/yaml/v2: v2.4.2
- go.yaml.in/yaml/v3: v3.0.4
- sigs.k8s.io/structured-merge-diff/v6: v6.3.0

### Changed
- github.com/emicklei/go-restful/v3: [v3.12.1 → v3.12.2](https://github.com/emicklei/go-restful/v3/compare/v3.12.1...v3.12.2)
- github.com/fxamacker/cbor/v2: [v2.7.0 → v2.9.0](https://github.com/fxamacker/cbor/v2/compare/v2.7.0...v2.9.0)
- github.com/google/gnostic-models: [v0.6.9 → v0.7.0](https://github.com/google/gnostic-models/compare/v0.6.9...v0.7.0)
- github.com/grpc-ecosystem/grpc-gateway/v2: [v2.24.0 → v2.26.3](https://github.com/grpc-ecosystem/grpc-gateway/v2/compare/v2.24.0...v2.26.3)
- github.com/modern-go/reflect2: [v1.0.2 → 35a7c28](https://github.com/modern-go/reflect2/compare/v1.0.2...35a7c28)
- github.com/spf13/cobra: [v1.8.1 → v1.9.1](https://github.com/spf13/cobra/compare/v1.8.1...v1.9.1)
- github.com/spf13/pflag: [v1.0.5 → v1.0.6](https://github.com/spf13/pflag/compare/v1.0.5...v1.0.6)
- go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc: v1.33.0 → v1.34.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace: v1.33.0 → v1.34.0
- go.opentelemetry.io/otel/metric: v1.33.0 → v1.35.0
- go.opentelemetry.io/otel/sdk: v1.33.0 → v1.34.0
- go.opentelemetry.io/otel/trace: v1.33.0 → v1.35.0
- go.opentelemetry.io/otel: v1.33.0 → v1.35.0
- go.opentelemetry.io/proto/otlp: v1.4.0 → v1.5.0
- google.golang.org/genproto/googleapis/api: e6fa225 → a0af3ef
- google.golang.org/genproto/googleapis/rpc: 9240e9c → a0af3ef
- google.golang.org/grpc: v1.69.0 → v1.72.1
- k8s.io/api: v0.33.0 → v0.34.0
- k8s.io/apimachinery: v0.33.0 → v0.34.1
- k8s.io/client-go: v0.33.0 → v0.34.0
- k8s.io/component-base: v0.33.0 → v0.34.0
- k8s.io/gengo/v2: a7b603a → 85fd79d
- k8s.io/kube-openapi: c8a335a → f3f2b99
- k8s.io/utils: 24370be → 4c0f3b2
- sigs.k8s.io/structured-merge-diff/v4: v4.6.0 → v4.5.0
- sigs.k8s.io/yaml: v1.4.0 → v1.6.0

### Removed
_Nothing has changed._
