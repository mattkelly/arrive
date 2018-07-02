SHELL=/bin/bash
PROJECT_NAME := "arrive"
PKG_LIST := $(shell glide novendor)

lint:
	@golint -set_exit_status ${PKG_LIST}

test:
	@go test ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

gen:
	@./hack/update-codegen.sh

verify:
	@./hack/verify-codegen.sh
