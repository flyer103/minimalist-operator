GODIR = $(shell go list ./... | grep -v /vendor/)
PKG := github.com/flyer103/minimalist-operator
BUILD_IMAGE ?= golang:1.19
GOARCH := amd64
# GOOS := linux
GOOS := darwin
VERSION := $(shell git rev-parse HEAD)

.PHONY: pre-build
pre-build:
	@echo "pre build"
	@echo "clean all flycheck files"
	@find . -name "flycheck*" | xargs rm -f

.PHONY: build-dirs
build-dirs: pre-build
	@mkdir -p .go/src/$(PKG) .go/bin .cache
	@mkdir -p release

.PHONY: build-resource
build-resource: pre-build
	@./hack/generate-groups.sh "deepcopy,client,informer,lister" \
	github.com/flyer103/minimalist-operator/pkg/clients \
	github.com/flyer103/minimalist-operator/pkg/apis \
	"example:v1alpha1"

.PHONY: build-operator
build-operator: build-dirs build-resource
	@docker run                                                            \
	    --rm                                                               \
	    -ti                                                                \
	    -u $$(id -u):$$(id -g)                                             \
	    -v $$(pwd)/.go:/go                                                 \
	    -v $$(pwd):/go/src/$(PKG)                                          \
	    -v $$(pwd)/release:/go/bin                                         \
	    -v $$(pwd)/.cache:/.cache            			                         \
	    -e GOOS=$(GOOS)                                                    \
	    -e GOARCH=$(GOARCH)                                                \
	    -e CGO_ENABLED=0                                                   \
	    -w /go/src/$(PKG)                                                  \
	    $(BUILD_IMAGE)                                                     \
	    go build -o ./release/operator ./cmd/
