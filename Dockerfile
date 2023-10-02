FROM --platform=$BUILDPLATFORM qmcgaw/xcputranslate:v0.6.0 AS xcputranslate

FROM --platform=$BUILDPLATFORM golang:1.21.1-alpine3.18 AS builder
COPY --from=xcputranslate /xcputranslate /usr/local/bin/xcputranslate

LABEL stage=gobuilder

ENV CGO_ENABLED 0

WORKDIR /build

COPY go.mod .
COPY *.go .
ARG TARGETOS
ARG TARGETARCH
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOOS=$TARGETOS \
    GOARCH="$(xcputranslate translate -targetplatform ${TARGETPLATFORM}  -language golang -field arch)" \
    GOARM="$(xcputranslate translate -targetplatform ${TARGETPLATFORM} -language golang -field arm)" \
    go build -ldflags="-s -w" -o /app/main main.go


FROM scratch

WORKDIR /app
COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["./main"]
