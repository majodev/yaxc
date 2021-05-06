all: init ##- Runs all of our common make targets: , init, build and test.
	go fmt ./...
	golangci-lint run --fast --timeout 5m
	@$(MAKE) test

test: ##- Run tests, output by package, print coverage.
	gotestsum --format pkgname-and-test-fails --jsonfile /tmp/test.log -- -race -cover -count=1 -coverprofile=/tmp/coverage.out ./...
	@printf "coverage "
	@go tool cover -func=/tmp/coverage.out | tail -n 1 | awk '{$$1=$$1;print}'

watch-tests: ##- Watches *.go files and runs package tests on modifications.
	gotestsum --format testname --watch -- -race -count=1

init: ##- Runs make modules, tools and tidy.
	go mod download
	go mod tidy

.PHONY: test

.ONESHELL:
SHELL = /bin/bash
.SHELLFLAGS = -ec
