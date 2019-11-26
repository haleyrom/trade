GO_ON ?= GO111MODULE=on go
GO_OFF ?= GO111MODULE=off go
GO ?= $(GO_ON)
GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell GO111MODULE=off $(GO) list ./...)
VETPACKAGES ?= $(shell GO111MODULE=off $(GO) list ./...)
GOFILES := $(shell find . -name "*.go" -type f ! -path "./vendor/*")

TAGS = jsoniter
TAGS_PPROF = $(TAGS) pprof

LDFLAGS += -X "weiliao/pkg/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "weiliao/pkg/version.GitHash=$(shell git rev-parse HEAD)"

.PHONY: default
default: build

.PHONY: ci
ci: lint vet test

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: tools
tools:
	$(GO_OFF) get golang.org/x/lint/golint
	$(GO_OFF) get github.com/client9/misspell/cmd/misspell