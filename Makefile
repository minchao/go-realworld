SHELL := /bin/bash -o pipefail

BUILD_VERSION ?= $(shell git describe --always)
BUILD_COMMIT ?= $(shell git rev-parse HEAD)
CMD_PACKAGE := github.com/minchao/go-realworld/cmd/realworld/cmd
DOCKER_IMAGE ?= realworld
DOCKER_VERSION ?= $(BUILD_VERSION)
DOCKERFILE ?= Dockerfile

.PHONY: help
## help: print this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: lint
## lint: run linters
lint:
	golangci-lint run ./...

.PHONY: test
## test: run unit tests
test:
	go test -v -race -coverprofile=coverage.out ./...

.PHONY: cover
## cover: open coverage profile
cover:
	go tool cover -html=coverage.out

.PHONY: build
## build: build app
build:
	CGO_ENABLED=0 go build \
		-ldflags "-s -X $(CMD_PACKAGE).Version=$(BUILD_VERSION) -X $(CMD_PACKAGE).Commit=$(BUILD_COMMIT) -X $(CMD_PACKAGE).Date=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'`" \
		./cmd/realworld

.PHONY: e2e
## e2e: end-to-end testing
e2e:
	newman run api/Conduit.postman_collection.json \
		-e api/e2e.postman_environment.json

.PHONY: docker-image
## docker-image: build docker image
docker-image:
	docker build \
		--build-arg BUILD_VERSION=$(BUILD_VERSION) \
		--build-arg BUILD_DATE=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` \
		--build-arg BUILD_COMMIT=$(BUILD_COMMIT) \
		--build-arg CMD_PACKAGE=$(CMD_PACKAGE) \
		-t $(DOCKER_IMAGE):$(DOCKER_VERSION) -f $(DOCKERFILE) .
