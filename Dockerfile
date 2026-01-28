FROM golang:1.25.6-alpine3.23 as builder

WORKDIR /go/src/github.com/mxssl/wait-for-pg
COPY . .

# Install external dependcies
RUN apk add --no-cache \
	ca-certificates=20251003-r0 \
	curl=8.17.0-r1 \
	git=2.52.0-r0

# Compile binary
RUN CGO_ENABLED=0 \
	GOOS="$(go env GOHOSTOS)" \
	GOARCH="$(go env GOHOSTARCH)" \
	go build -v -o wait-for-pg

# Copy compiled binary to clear Alpine Linux image
FROM alpine:3.23
WORKDIR /
RUN apk add --no-cache ca-certificates=20251003-r0
COPY --from=builder /go/src/github.com/mxssl/wait-for-pg/wait-for-pg /usr/local/bin/wait-for-pg
RUN chmod +x /usr/local/bin/wait-for-pg
