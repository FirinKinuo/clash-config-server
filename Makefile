PROJECT := clash-config-server
BUILD_TIME ?= $(shell date +%d.%m.%Y-%H%M)
GIT_COMMIT_HASH := $(shell git rev-parse --short HEAD)

VERSION ?= $(GIT_COMMIT_HASH)@$(BUILD_TIME)
BUILD_NAME ?= $(PROJECT)-$(BUILD_TIME)

GO_MAIN := ""

RELEASE_LDFLAGS := "-s -w -X main.version=$(VERSION)"
RELEASE_BUILD := go build -ldflags $(RELEASE_LDFLAGS) -v

DOCKER_LD_FLAGS := "-s -w -X main.version=docker-$(GIT_COMMIT_HASH)"
DOCKER_BUILD := go build -ldflags $(DOCKER_LD_FLAGS) -v

DEV_LDFLAGS := "-X main.version=dev@$(BUILD_TIME)-$(GIT_COMMIT_HASH)"
DEV_BUILD := go build -ldflags $(DEV_LDFLAGS) -v

.PHONY: clean
clean:
	rm -rf _build/ release/

.PHONY: build
build:
	$(RELEASE_BUILD) -o $(BUILD_NAME) $(GO_MAIN)

build-docker:
	$(DOCKER_BUILD) -o $(PROJECT) $(GO_MAIN)

.PHONY: build-dev
build-dev:
	$(DEV_BUILD) -o $(PROJECT) $(GO_MAIN)

.PHONY: build-all
build-all:
	- make clean
	mkdir _build
	GOOS=windows GOARCH=amd64 $(RELEASE_BUILD) -o _build/$(PROJECT)-windows-amd64.exe $(GO_MAIN)
	GOOS=windows GOARCH=arm64 $(RELEASE_BUILD) -o _build/$(PROJECT)-windows-arm64.exe $(GO_MAIN)
	GOOS=linux GOARCH=amd64 $(RELEASE_BUILD) -o _build/$(PROJECT)-linux-amd64.exe $(GO_MAIN)
	GOOS=linux GOARCH=arm64 $(RELEASE_BUILD) -o _build/$(PROJECT)-linux-arm64.exe $(GO_MAIN)
	cd _build; sha256sum * > sha256sums.txt

.PHONY: release
release: clean build-all
	mkdir release
	cp _build/* release
	cd release; sha256sum --quiet --check sha256sums.txt