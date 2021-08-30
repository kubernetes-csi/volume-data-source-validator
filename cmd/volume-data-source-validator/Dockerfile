FROM gcr.io/distroless/base:latest
LABEL maintainers="Kubernetes Authors"
LABEL description="Volume Data Source Validator Controller"
ARG binary=./bin/volume-data-source-validator

COPY ${binary} volume-data-source-validator
ENTRYPOINT ["/volume-data-source-validator"]
