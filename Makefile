SHELL := /bin/bash -o pipefail

BUILD_VERSION ?= $(shell git describe --always)
BUILD_COMMIT ?= $(shell git rev-parse HEAD)
CMD_PACKAGE := github.com/minchao/go-realworld/cmd/realworld/cmd

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
	go build \
		-ldflags "-s -X $(CMD_PACKAGE).Version=$(BUILD_VERSION) -X $(CMD_PACKAGE).Date=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X $(CMD_PACKAGE).Commit=$(BUILD_COMMIT)" \
		./cmd/realworld
