# Default bundle image tag
BUNDLE_IMG ?= controller-bundle:$(VERSION)
# Options for 'bundle-build'
ifneq ($(origin CHANNELS), undefined)
BUNDLE_CHANNELS := --channels=$(CHANNELS)
endif
ifneq ($(origin DEFAULT_CHANNEL), undefined)
BUNDLE_DEFAULT_CHANNEL := --default-channel=$(DEFAULT_CHANNEL)
endif
BUNDLE_METADATA_OPTS ?= $(BUNDLE_CHANNELS) $(BUNDLE_DEFAULT_CHANNEL)

# Image URL to use all building/pushing image targets
IMG ?= controller:latest
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"
DBG         ?= 0
PROJECT     ?= federatorai-operator
ORG_PATH    ?= github.com/containers-ai
REPO_PATH   ?= $(ORG_PATH)/$(PROJECT)
VERSION     = $(shell git rev-parse --abbrev-ref HEAD)-$(shell git rev-parse --short HEAD)$(TMP_VERSION_SUFFIX)$(shell git diff --quiet || echo '-dirty')
LD_FLAGS    ?= -X $(REPO_PATH)/pkg/version.Raw=$(VERSION)
BUILD_DEST  ?= bin/federatorai-operator
MUTABLE_TAG ?= latest
IMAGE		= federatorai-operator
IMAGE_TAG	?= $(shell git describe --always --dirty --abbrev=7)

FIRST_GOPATH:=$(firstword $(subst :, ,$(shell go env GOPATH)))
GOBINDATA_BIN=$(FIRST_GOPATH)/bin/go-bindata

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

ifeq ($(DBG),1)
GOGCFLAGS ?= -gcflags=all="-N -l"
endif

.PHONY: all
all: build images check

NO_DOCKER ?= 0
ifeq ($(NO_DOCKER), 1)
  DOCKER_CMD =
  IMAGE_BUILD_CMD = imagebuilder
else
  DOCKER_CMD := docker run --rm -v "$(PWD):/go/src/$(REPO_PATH):Z" -w "/go/src/$(REPO_PATH)" openshift/origin-release:golang-1.13
  IMAGE_BUILD_CMD = docker build
endif

.PHONY: pkg/assets/bindata.go
pkg/assets/bindata.go: $(GOBINDATA_BIN)
	# Using "-modtime 1" to make generate target deterministic. It sets all file time stamps to unix timestamp 1
	cd assets && $(GOBINDATA_BIN) -pkg assets -o ../$@ \
		./... \

.PHONY: pkg/patch/assets/bindata.go
pkg/patch/assets/bindata.go: $(GOBINDATA_BIN)
	cd pkg/patch/assets && $(GOBINDATA_BIN) -pkg assets -o ../../../$@ \
		./... \

.PHONY: build
build: pkg/assets/bindata.go pkg/patch/assets/bindata.go crd-check generate ## build binaries
	$(DOCKER_CMD) go build -mod=vendor $(GOGCFLAGS) -ldflags "$(LD_FLAGS)" -o "$(BUILD_DEST)" "$(REPO_PATH)"

.PHONY: images
images: ## Create images
	$(IMAGE_BUILD_CMD) -t "$(IMAGE):$(IMAGE_TAG)" -t "$(IMAGE):$(MUTABLE_TAG)" -f Dockerfile.ubi ./

.PHONY: push
push:
	docker push "$(IMAGE):$(IMAGE_TAG)"
	docker push "$(IMAGE):$(MUTABLE_TAG)"

.PHONY: check
check: fmt vet test ## Check your code

# Run tests
.PHONY: test
test: generate fmt vet manifests
	go test ./... -coverprofile cover.out	

# Run go fmt against code
.PHONY: fmt
fmt:
	go fmt ./...	

# Run go vet against code
.PHONY: vet
vet:
	go vet ./...	

$(GOBINDATA_BIN):
	go get -mod=readonly -u github.com/shuLhan/go-bindata/...

.PHONY: help
help:
	@grep -E '^[a-zA-Z/0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: crd-check
crd-check: manifests
	if [ "" != "`diff -du config/crd/bases/federatorai.containers.ai_alamedaservices.yaml deploy/upstream/02-alamedaservice.crd.yaml 2>&1`" ]; then \
	    echo "Error. Improper CRDs files."; \
	    exit 1; \
	fi

# Install CRDs into a cluster
.PHONY: install
install: manifests kustomize
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
.PHONY: uninstall
uninstall: manifests kustomize
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
.PHONY: deploy
deploy: manifests kustomize
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
.PHONY: manifests
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: kustomize
kustomize:
ifeq (, $(shell which kustomize))
	@{ \
	set -e ;\
	KUSTOMIZE_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$KUSTOMIZE_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/kustomize/kustomize/v3@v3.5.4 ;\
	rm -rf $$KUSTOMIZE_GEN_TMP_DIR ;\
	}
KUSTOMIZE=$(GOBIN)/kustomize
else
KUSTOMIZE=$(shell which kustomize)
endif

# Generate bundle manifests and metadata, then validate generated files.
.PHONY: bundle
bundle: manifests
	operator-sdk generate kustomize manifests -q
	kustomize build config/manifests | operator-sdk generate bundle -q --overwrite --version $(VERSION) $(BUNDLE_METADATA_OPTS)
	operator-sdk bundle validate ./bundle

# Build the bundle image.
.PHONY: bundle-build
bundle-build:
	docker build -f bundle.Dockerfile -t $(BUNDLE_IMG) .

# Generate code
.PHONY: generate
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

# Build the docker image
.PHONY: docker-build
docker-build: test
	docker build . -t ${IMG}

# Push the docker image
.PHONY: docker-push
docker-push:
	docker push ${IMG}

# find or download controller-gen
# download controller-gen if necessary
.PHONY: controller-gen
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get -mod=readonly sigs.k8s.io/controller-tools/cmd/controller-gen@v0.3.0 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	go run ./main.go