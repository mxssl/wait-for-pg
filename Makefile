BINARY_NAME=wait-for-pg
CURRENT_DIR=$(shell pwd)
TAG=$(shell git name-rev --tags --name-only $(shell git rev-parse HEAD))
DOCKER_REGISTRY=mxssl
export GO111MODULE=on

.PHONY: all build clean lint test init github-release-dry github-release docker-release

all: build

build:
	go build -v -mod=vendor -o ${BINARY_NAME}

clean:
	rm -f ${BINARY_NAME}

lint:
	golangci-lint run -v

test:
	go test -v ./...

init:
	go mod init

tidy:
	go mod tidy

github-release-dry:
	@echo "TAG: ${TAG}"
	goreleaser release --rm-dist --snapshot --skip-publish

github-release:
	@echo "TAG: ${TAG}"
	goreleaser release --rm-dist

docker-release:
	@echo "Registry: ${DOCKER_REGISTRY}"
	@echo "TAG: ${TAG}"
	docker build --tag ${DOCKER_REGISTRY}/${BINARY_NAME}:${TAG} .
	docker push ${DOCKER_REGISTRY}/${BINARY_NAME}:${TAG}
