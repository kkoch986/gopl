BIN_DIR=bin
BIN_NAME=gopl
PACKAGE_NAME=$(shell go list .)
SOURCE_DIR=.
SOURCES=$(shell find $(SOURCE_DIR) -name '*.go')
GO111MODULE?=on

all: build

build: parser $(BIN_DIR)/$(BIN_NAME)

test: build 
	go test ./...

parser: grammar.md
	gogll grammar.md

fmt:
	go fmt ./...

check-fmt:
	@gofmt -d . | read; if [ $$? == 0 ]; then echo "gofmt check failed for:"; gofmt -d -l .; exit 1; fi

$(BIN_DIR)/$(BIN_NAME): $(SOURCES)
	mkdir -p $(BIN_DIR)
	GO111MODULE=$(GO111MODULE) \
	CGO_ENABLED=1 \
		go build -x -o $(BIN_DIR)/$(BIN_NAME) -ldflags "\
		-linkmode internal \
		-extldflags '-static' \
		-v \
		-s -w \
		-X $(PACKAGE_NAME)/version.BuildTime=$(shell date -u +%FT%T%z)\
		-X $(PACKAGE_NAME)/version.GitCommit=$(shell git rev-parse --short HEAD)\
		-X $(PACKAGE_NAME)/version.Version=$(shell git describe --abbrev=0 --tags 2> /dev/null || echo v0.0.1)"

clean:
	rm -rf $(BIN_DIR)

go-lint:
	$(eval GOLINT_INSTALLED := $(shell which golangci-lint))

	@if [ "$(GOLINT_INSTALLED)" = "" ]; then \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.16.0; \
	fi;

lint: go-lint
	golangci-lint run
