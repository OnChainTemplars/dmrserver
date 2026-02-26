BUILDDIR ?= $(shell pwd)/build/

.PHONY: fmt
fmt:
	gofmt -s -w ./.

.PHONY: build
build:
	go build -o $(BUILDDIR) ./cmd/dmrserver

.PHONY: run
run:
	go run ./cmd/dmrserver
