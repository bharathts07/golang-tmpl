GIT_REF := $(shell git describe --always)
VERSION ?= commit-$(GIT_REF)

SERVICE_NAME := $(shell grep "^module" go.mod | rev | cut -d "/" -f1 | rev)


GCR_PROJECT := pokke
REGISTRY := gcr.io/$(GCR_PROJECT)
# For use in cloudbuild if necessary
GCR_IMAGE_NAME := $(REGISTRY)/$(SERVICE_NAME):$(VERSION)
# For pushing built docker images to public docker hub
DOCKER_HUB_IMAGE := bharathts07/$(SERVICE_NAME):$(VERSION)

.PHONY: build
build: build/server

build/server:
	CGO_ENABLED=0 go build -o bin/server \
	-ldflags "-X main.version=$(VERSION) -X main.serviceName=$(SERVICE_NAME)" \
	./cmd/server

.PHONY: test
test: unit-test

.PHONY: unit-test
unit-test:
	@go test -count=1 -race -v $(shell go list ./...)

.PHONY: mockgenerate
mockgenerate:
	@echo "Generating mock for HomeData update service"
	mockgen -source=./database/interface.go -destination=./database/mock.go

# GITHUB_TOKEN needed if the image needs to be pulled from a private repository
.PHONY: container
container:
	@docker build -t $(DOCKER_HUB_IMAGE) \
		--build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) \
		--build-arg VERSION=$(VERSION) \
		--build-arg SERVICE_NAME=$(SERVICE_NAME) \
		--file docker/server/Dockerfile \
		.

# This requires `docker login`
.PHONY: push-container
push-container:
	@docker push $(DOCKER_HUB_IMAGE)
