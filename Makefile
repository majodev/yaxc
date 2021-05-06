### -----------------------
# --- Building
### -----------------------

# first is default target when running "make" without args
build: ##- Default 'make' target: sql, swagger, go-generate, go-format, go-build and lint.
	@$(MAKE) go-format
	@$(MAKE) go-build

# useful to ensure that everything gets resetuped from scratch
all:  init ##- Runs all of our common make targets: , init, build and test.
	@$(MAKE) build
	@$(MAKE) test

go-format: ##- (opt) Runs go format.
	go fmt ./...

go-build: ##- (opt) Runs go build.
	go build -o bin/app

go-lint: ##- (opt) Runs golangci-lint.
	golangci-lint run --fast --timeout 5m

# https://github.com/gotestyourself/gotestsum#format 
# w/o cache https://github.com/golang/go/issues/24573 - see "go help testflag"
# note that these tests should not run verbose by default (e.g. use your IDE for this)
# TODO: add test shuffling/seeding when landed in go v1.15 (https://github.com/golang/go/issues/28592)
# tests by pkgname
test: ##- Run tests, output by package, print coverage.
	@$(MAKE) go-test-by-pkg
	@$(MAKE) go-test-print-coverage

# note that we explicitly don't want to use a -coverpkg=./... option, per pkg coverage take precedence
go-test-by-pkg: ##- (opt) Run tests, output by package.
	gotestsum --format pkgname-and-test-fails --jsonfile /tmp/test.log -- -race -cover -count=1 -coverprofile=/tmp/coverage.out ./...

go-test-print-coverage: ##- (opt) Print overall test coverage (must be done after running tests).
	@printf "coverage "
	@go tool cover -func=/tmp/coverage.out | tail -n 1 | awk '{$$1=$$1;print}'

watch-tests: ##- Watches *.go files and runs package tests on modifications.
	gotestsum --format testname --watch -- -race -count=1

### -----------------------
# --- Initializing
### -----------------------

init: ##- Runs make modules, tools and tidy.
	@$(MAKE) modules
	@$(MAKE) tidy

# cache go modules (locally into .pkg)
modules: ##- (opt) Cache packages as specified in go.mod.
	go mod download

tidy: ##- (opt) Tidy our go.sum file.
	go mod tidy

# https://www.gnu.org/software/make/manual/html_node/Special-Targets.html
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
# ignore matching file/make rule combinations in working-dir
.PHONY: test

# https://unix.stackexchange.com/questions/153763/dont-stop-makeing-if-a-command-fails-but-check-exit-status
# https://www.gnu.org/software/make/manual/html_node/One-Shell.html
# required to ensure make fails if one recipe fails (even on parallel jobs)
.ONESHELL:
SHELL = /bin/bash
.SHELLFLAGS = -ec
