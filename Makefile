PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)
PKG_NAME := "warehouse"
GIT_COMMIT:=$(shell git rev-parse --verify HEAD --short=7)
GO_VERSION:=$(shell go version | grep -o "go1\.[0-9|\.]*")
VERSION?=0.0.0
BIN_NAME=warehouse
APP_NAME?=warehouse
APP_NAME_UPPER=`echo $(APP_NAME) | tr '[:lower:]' '[:upper:]'`

binary: clean fmt
	@echo "Building binary for commit $(GIT_COMMIT)"
	go build -ldflags="-X github.com/junland/warehouse/cmd.BinVersion=$(VERSION) -X github.com/junland/warehouse/cmd.GoVersion=$(GO_VERSION)" -o $(BIN_NAME)

.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -f ./warehouse
	@rm -rf ./warehouse-*
	@rm -rf ./*.tar.gz
	@rm -rf ./warehouse_*
	@rm -rf ./*.txt
	@rm -rf ./*.pem
	@echo "Done cleaning..."

.PHONY: fmt
fmt:
	@echo "Running $@"
	go fmt ./...

.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

.PHONY: binary-race
binary-race: clean
	@echo "Building race binary for commit $(GIT_COMMIT)"
	go build -race -ldflags="-X github.com/junland/warehouse/cmd.BinVersion=$(VERSION) -X github.com/junland/warehouse/cmd.GoVersion=$(GO_VERSION)" -o $(BIN_NAME)

.PHONY: tls-certs
tls-certs:
	@echo "Making Development TLS Certificates..."
	@openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -subj "/C=US/ST=Texas/L=Austin/O=Local Development/OU=IT Department/CN=127.0.0.0"

.PHONY: travis-sizes
travis-sizes: clean fmt
	@echo "Building unstripped binary..."
	@go build -o warehouse-default || (echo "Failed to build binary: $$?"; exit 1)
	@echo "Size of unstripped binary: $$(ls -l warehouse-default | awk '{print $$5}') bytes or $$(ls -lh warehouse-default | awk '{print $$5}')" > ./size-report.txt
	@echo "Building stripped binary..."
	@go build -ldflags="-s -w" -o warehouse-stripped || (echo "Failed to build stripped binary: $$?"; exit 1)
	@echo "Size of stripped binary: $$(ls -l warehouse-stripped | awk '{print $$5}') bytes or $$(ls -lh warehouse-stripped | awk '{print $$5}')" >> ./size-report.txt
	@echo "Compressing stripped binary..."
	@cp ./warehouse-stripped ./warehouse-compressed
	@upx -9 -q ./warehouse-compressed > /dev/null || (echo "Failed to compress stripped binary: $$?"; exit 1)
	@echo "Size of compressed stripped binary: $$(ls -l warehouse-compressed | awk '{print $$5}') bytes or $$(ls -lh warehouse-compressed | awk '{print $$5}')" >> ./size-report.txt
	@echo "Building binary with gccgo..."
	@go build -compiler gccgo -o warehouse-gccgo
	@echo "Size of binary using gccgo: $$(ls -l warehouse-gccgo | awk '{print $$5}') bytes or $$(ls -lh warehouse-gccgo | awk '{print $$5}')" >> ./size-report.txt
	@echo "Reported binary sizes for Go version $$(echo -n $$(go version) | grep -o '1\.[0-9|\.]*'): "
	@cat ./size-report.txt
	@rm -f ./*.txt

.PHONY: travis-sizes-nogccgo
travis-sizes-nogccgo: clean fmt
	@echo "Building unstripped binary..."
	@go build -o warehouse-default || (echo "Failed to build binary: $$?"; exit 1)
	@echo "Size of unstripped binary: $$(ls -l warehouse-default | awk '{print $$5}') bytes or $$(ls -lh warehouse-default | awk '{print $$5}')" > ./size-report.txt
	@echo "Building stripped binary..."
	@go build -ldflags="-s -w" -o warehouse-stripped || (echo "Failed to build stripped binary: $$?"; exit 1)
	@echo "Size of stripped binary: $$(ls -l warehouse-stripped | awk '{print $$5}') bytes or $$(ls -lh warehouse-stripped | awk '{print $$5}')" >> ./size-report.txt
	@echo "Compressing stripped binary..."
	@cp ./warehouse-stripped ./warehouse-compressed
	@upx -9 -q ./warehouse-compressed > /dev/null || (echo "Failed to compress stripped binary: $$?"; exit 1)
	@echo "Size of compressed stripped binary: $$(ls -l warehouse-compressed | awk '{print $$5}') bytes or $$(ls -lh warehouse-compressed | awk '{print $$5}')" >> ./size-report.txt
	@echo "Reported binary sizes for Go version $$(echo -n $$(go version) | grep -o '1\.[0-9|\.]*'): "
	@cat ./size-report.txt
	@rm -f ./*.txt

.PHONY: slim-binary
slim-binary: clean fmt
	rm -f $(BIN_NAME)
	go build -ldflags "-s -w -X github.com/junland/warehouse/cmd.BinVersion=$(VERSION) -X github.com/junland/warehouse/cmd.GoVersion=$(GO_VERSION)" -o $(BIN_NAME)

.PHONY: amd64-binary
amd64-binary: clean fmt
	rm -f $(BIN_NAME)
	GOARCH=amd64 go build -ldflags "-X github.com/junland/warehouse/cmd.BinVersion=$(VERSION) -X github.com/junland/warehouse/cmd.GoVersion=$(GO_VERSION)" -o $(BIN_NAME)

.PHONY: aarch64-binary
aarch64-binary: clean fmt
	rm -f $(BIN_NAME)
	GOARCH=arm64 go build -ldflags "-X github.com/junland/warehouse/cmd.BinVersion=$(VERSION) -X github.com/junland/warehouse/cmd.GoVersion=$(GO_VERSION)" -o $(BIN_NAME)

.PHONY: armhf-binary
armhf-binary: clean fmt
	rm -f $(BIN_NAME)
	GOARCH=arm GOARM=7 go build -ldflags "-X github.com/junland/warehouse/cmd.BinVersion=$(VERSION) -X github.com/junland/warehouse/cmd.GoVersion=$(GO_VERSION)" -o $(BIN_NAME)