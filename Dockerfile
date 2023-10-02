FROM golang:1.21.1-alpine3.18 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

WORKDIR /build

COPY go.mod .
COPY *.go .
RUN go build -ldflags="-s -w" -o /app/main main.go


FROM scratch

WORKDIR /app
COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["./main"]
