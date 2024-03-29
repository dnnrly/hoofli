GO111MODULE=on

CURL_BIN ?= curl
GO_BIN ?= go
GORELEASER_BIN ?= goreleaser

PUBLISH_PARAM?=
GO_MOD_PARAM?=-mod vendor
TMP_DIR?=./tmp

BASE_DIR=$(shell pwd)

NAME=hoofli

export GO111MODULE=on
export GOPROXY=https://proxy.golang.org
export PATH := $(BASE_DIR)/bin:$(PATH)

.PHONY: install deps clean clean-deps test-deps build-deps deps test acceptance-test ci-test lint release update

install: ## install hoofli from the current working tree
	$(GO_BIN) install -v .

build: ## build hoofli
	$(GO_BIN) build -v -o ./$(NAME) ./cmd/$(NAME)

clean: ## remove build artifacts from the working tree
	rm -f $(NAME)
	rm -rf dist/
	rm -rf cmd/$(NAME)/dist

clean-deps: ## remove dependencies in the working tree
	rm -rf ./bin
	rm -rf ./tmp
	rm -rf ./libexec
	rm -rf ./share

./bin/golangci-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.37.1

./bin/tparse: ./bin ./tmp
	curl --fail -L -o ./tmp/tparse.tar.gz https://github.com/mfridman/tparse/releases/download/v0.8.3/tparse_0.8.3_Linux_x86_64.tar.gz
	tar -xf ./tmp/tparse.tar.gz -C ./bin

./bin/godog: ./bin ./tmp
	curl --fail -L -o ./tmp/godog.tar.gz https://github.com/cucumber/godog/releases/download/v0.11.0/godog-v0.11.0-linux-amd64.tar.gz
	tar -xf ./tmp/godog.tar.gz -C ./tmp
	cp ./tmp/godog-v0.11.0-linux-amd64/godog ./bin


test-deps: ./bin/godog ./bin/tparse ./bin/golangci-lint ## ci target - install test dependencies
	$(GO_BIN) get -v ./...
	$(GO_BIN) mod tidy

./bin:
	mkdir ./bin

./tmp:
	mkdir ./tmp

./bin/goreleaser: ./bin ./tmp
	$(CURL_BIN) --fail -L -o ./tmp/goreleaser.tar.gz https://github.com/goreleaser/goreleaser/releases/download/v0.168.2/goreleaser_Linux_x86_64.tar.gz
	gunzip -f ./tmp/goreleaser.tar.gz
	tar -C ./bin -xvf ./tmp/goreleaser.tar

build-deps: ./bin/goreleaser ## ci target - install build dependencies

deps: build-deps test-deps ## ci target - install build and test dependencies

test: ## run unit tests with tparse prettifying
	$(GO_BIN) test -json ./... | tparse -all

acceptance-test: ## run acceptance tests on built hoofli
	cd test && godog -t @Acceptance
 
ci-test: ## ci target - run unit tests
	$(GO_BIN) test -race -coverprofile=coverage.txt -covermode=atomic ./...

lint: ## run linting
	golangci-lint run

release: clean ## ci target - release hoofli
	cd cmd/$(NAME) ; $(GORELEASER_BIN) $(PUBLISH_PARAM)

update: ## update dependencies
	$(GO_BIN) get -u
	$(GO_BIN) mod tidy
	make test
	make install
	$(GO_BIN) mod tidy

help:   ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/:.\+##/ --/'
