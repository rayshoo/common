ifneq (,$(wildcard ./.env))
	include .env
	export
endif
GIT_VER=$(shell git describe --tags 2>/dev/null)
.DEFAULT_GOAL := build

imports:
	goimports -l -w .
.PHONY:imports

fmt: imports
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	CGO_ENABLED=0 go vet ./...
.PHONY:vet

build: vet
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/${COMMAND}-linux-amd64 -ldflags "-s -w -X main.command=${COMMAND} -X main.version=${GIT_VER}"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/${COMMAND}-windows-amd64 -ldflags "-s -w -X main.command=${COMMAND} -X main.version=${GIT_VER}"
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/${COMMAND}-darwin-amd64 -ldflags "-s -w -X main.command=${COMMAND} -X main.version=${GIT_VER}"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/${COMMAND}-linux-arm64 -ldflags "-s -w -X main.command=${COMMAND} -X main.version=${GIT_VER}"
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o build/${COMMAND}-windows-arm64 -ldflags "-s -w -X main.command=${COMMAND} -X main.version=${GIT_VER}"
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/${COMMAND}-darwin-arm64 -ldflags "-s -w -X main.command=${COMMAND} -X main.version=${GIT_VER}"
.PHONY:build

install:
	CGO_ENABLED=0 go install -ldflags "-s -w -X main.command=${COMMAND} -X main.version=${GIT_VER}"
.PHONY:install