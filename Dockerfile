FROM golang:1.19.3-alpine3.15 as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/mxssl/wait-for-pg
COPY . .

# Install external dependcies
RUN apk add --no-cache \
  ca-certificates \
  curl \
  git

# Compile binary
RUN CGO_ENABLED=0 \
  GOOS=`go env GOHOSTOS` \
  GOARCH=`go env GOHOSTARCH` \
  go build -v -o wait-for-pg

# Copy compiled binary to clear Alpine Linux image
FROM alpine:3.21.3
WORKDIR /
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/mxssl/wait-for-pg/wait-for-pg /usr/local/bin/wait-for-pg
RUN chmod +x /usr/local/bin/wait-for-pg
