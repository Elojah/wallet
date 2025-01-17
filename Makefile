PACKAGE   = wallet
DATE     ?= $(shell date +%FT%T%z)
VERSION  ?= $(shell echo $(shell cat $(PWD)/.version)-$(shell git describe --tags --always))

GO = go

ifneq ($(wildcard ./bin/golangci-lint),)
	GOLINT = ./bin/golangci-lint
else
	GOLINT = golangci-lint
endif

GODOC       = godoc
GOFMT       = gofmt

API         = api

V         = 0
Q         = $(if $(filter 1,$V),,@)
M         = $(shell printf "\033[0;35m▶\033[0m")

.PHONY: all

all: api

api:  ## Build api binary
	$(info $(M) building executable api…) @
	$Q cd cmd/$(API) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(API)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(API)_$(VERSION) bin/$(PACKAGE)_$(API)

# Utils
.PHONY: proto
proto: ## Generate .proto files
	$(info $(M) running protobuf…) @
	$Q cd pkg/wallet && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. tx.proto
	$Q cd pkg/wallet && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. wallet.proto

# Vendoring
.PHONY: vendor
vendor: ## Write dependencies into vendor
	$(info $(M) running go mod vendor…) @
	$Q $(GO) mod vendor

.PHONY: tidy
tidy: ## Remove unused dependencies and add new required
	$(info $(M) running go mod tidy…) @
	$Q $(GO) mod tidy

# Check
.PHONY: check ## lint + test
check: vendor lint test

# Lint
.PHONY: lint
lint: ## Check code respect linter rules
	$(info $(M) running $(GOLINT)…)
	$Q $(GOLINT) run --deadline=5m

# Test
.PHONY: test
test: ## Run unit tests only
	$(info $(M) running go test…) @
	$Q $(GO) test -cover -race -v ./...

# Helpers
go-version: ## Print go version used in this makefile
	$Q echo $(GO)

.PHONY: fmt
fmt: ## Format code
	$(info $(M) running $(GOFMT)…) @
	$Q $(GOFMT) ./...

.PHONY: doc
doc: ## Generate project documentation
	$(info $(M) running $(GODOC)…) @
	$Q $(GODOC) ./...

.PHONY: clean
clean: ## Clean generated binaries
	$(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf bin/$(PACKAGE)_*

.PHONY: version
version: ## Print current project version
	@echo $(VERSION)

.PHONY: help
help: ## Print this
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
