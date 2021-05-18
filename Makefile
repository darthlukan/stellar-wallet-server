BINARY := stellar-wallet-server
VERSION := v0.0.1
ARCH := amd64

.PHONY: all
all: build container

.PHONY: build
build:
	mkdir -p build
	GOOS=linux GOARCH=$(ARCH) go build -o build/$(BINARY) .


container: build
	podman build -f Dockerfile -t quay.io/btomlins/stellar-wallet-server:$(VERSION)

.PHONY: clean
clean:
	rm build/*
