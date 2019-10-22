GIT_REF := $(shell git describe --always)
VERSION ?= commit-$(GIT_REF)

SERVICE_NAME := $(shell grep "^module" go.mod | rev | cut -d "/" -f1 | rev)

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
