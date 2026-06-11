FROM golang:1.26.3-alpine3.23 AS builder

WORKDIR /go/src/github.com/mxssl/wait-for-pg
COPY . .

# Install external dependencies
RUN apk add --no-cache \
	ca-certificates=20260413-r0 \
	curl=8.19.0-r0 \
	git=2.52.0-r0

# Compile binary
RUN CGO_ENABLED=0 \
	GOOS="$(go env GOHOSTOS)" \
	GOARCH="$(go env GOHOSTARCH)" \
	go build -v -o wait-for-pg

# Copy compiled binary to clear Alpine Linux image
FROM alpine:3.24
WORKDIR /
RUN apk add --no-cache ca-certificates=20260413-r0
COPY --from=builder /go/src/github.com/mxssl/wait-for-pg/wait-for-pg /usr/local/bin/wait-for-pg
RUN chmod +x /usr/local/bin/wait-for-pg
