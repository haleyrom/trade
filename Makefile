GO_ON ?= GO111MODULE=on go
GO_OFF ?= GO111MODULE=off go
GO ?= $(GO_ON)
GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell GO111MODULE=off $(GO) list ./...)
VETPACKAGES ?= $(shell GO111MODULE=off $(GO) list ./...)
GOFILES := $(shell find . -name "*.go" -type f ! -path "./vendor/*")

TAGET_APP = receive_msg
TAGS = jsoniter
TAGS_PPROF = $(TAGS) pprof

LDFLAGS += -X "github.com/haleyrom/trade/pkg/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "ithub.com/haleyrom/trade/pkg/version.GitHash=$(shell git rev-parse HEAD)"

.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')

SERVER_BIN = "./cmd/server/server"
RELEASE_ROOT = "release"
RELEASE_SERVER = "release/server"


all: start

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

build:
	@$(GO) build -ldflags '$(LDFLAGS)' -tags '$(TAGS)' -o $(SERVER_BIN) ./cmd/server

start:
	@go run -ldflags $(SERVER_BIN) ./cmd/server
	$(SERVER_BIN) -c ./assets/config/conf.yaml

serve:
	@$(GO) run -ldflags '$(LDFLAGS)' -tags '$(TAGS)' ./cmd/server/server.go

test:
	@go test -cover -race ./...

clean:
	rm -rf data release $(SERVER_BIN) ./internal/app/test/data ./cmd/server/data

pack: build
	rm -rf $(RELEASE_ROOT)
	mkdir -p $(RELEASE_SERVER)
	cp -r $(SERVER_BIN) configs $(RELEASE_SERVER)
	cd $(RELEASE_ROOT) && zip -r server.$(NOW).zip "server"