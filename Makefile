BINARY := stellar-wallet-server
VERSION := v0.0.1
ARCH := amd64

.PHONY: all
all: test build container

.PHONY: test
test:
	go test -v ./stellar-wallet-server_test.go
	go test -v ./account
	go test -v ./api

.PHONY: build
build:
	mkdir -p build
	GOOS=linux GOARCH=$(ARCH) go build -o build/$(BINARY) .


container: build
	podman build -f Dockerfile -t quay.io/btomlins/stellar-wallet-server:$(VERSION)

.PHONY: clean
clean:
	rm build/*
