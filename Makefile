BINARY := stellar-wallet-server
REGISTRY := quay.io
REPO := btomlins
VERSION := v0.0.2
ARCH := amd64
OS := linux

.PHONY: all
all: test build container

.PHONY: test
test:
	go test -v ./harness
	go test -v ./stellar-wallet-server_test.go
	go test -v ./account
	go test -v ./api

.PHONY: build
build:
	mkdir -p build
	GOOS=$(OS) GOARCH=$(ARCH) go build -o build/$(BINARY) .

container: build
	podman build -f Dockerfile -t $(REGISTRY)/$(REPO)/$(BINARY):$(VERSION)
	podman tag $(REGISTRY)/$(REPO)/$(BINARY):$(VERSION) $(REGISTRY)/$(REPO)/$(BINARY):latest
	podman push $(REGISTRY)/$(REPO)/$(BINARY):$(VERSION)
	podman push $(REGISTRY)/$(REPO)/$(BINARY):latest

.PHONY: clean
clean:
	rm build/*
