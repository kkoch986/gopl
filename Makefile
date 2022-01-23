BIN_DIR=bin
BIN_NAME=gopl
PACKAGE_NAME=$(shell go list .)
PACKAGES=$(shell go list ./...)
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

go-lint:
	$(eval GOLINT_INSTALLED := $(shell which golangci-lint))
	@if [ "$(GOLINT_INSTALLED)" = "" ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.35.2; \
	fi;

lint: go-lint
	golangci-lint run


test-cov: acc
	GO111MODULE=$(GO111MODULE) go-acc -o coverage.txt $(PACKAGES)

cov-html: test-cov
	go tool cover -html=coverage.txt

cov-func: test-cov
	go tool cover -func=coverage.txt

html-cov-report:
	go tool cover -html=coverage.txt -o coverage.html

clean:
	rm -rf $(BIN_DIR)
	rm -f .coverage*

acc:
	$(eval GO_ACC_INSTALLED := $(shell which go-acc))

	@if [ "$(GO_ACC_INSTALLED)" = "" ]; then \
		GO111MODULE=on go get -u github.com/ory/go-acc; \
	fi;
