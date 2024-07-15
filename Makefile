include .env

PROJECTNAME=$(shell basename "$(PWD)")

# go environment variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
CLI_TOOL=github.com/binaryshogun/pdftoimage/cmd/cli

# writing errors into the file.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# writing PID into the file.
PID=/tmp/.$(PROJECTNAME)-api-server.pid

## Writes command descriptions.
.PHONY: help
all: help
help: Makefile
	@awk ' \
		/^## / { \
			help_msg = substr($$0, 4); \
			getline; \
			if (match($$0, /^([a-zA-Z_-]+):/)) { \
				cmd = substr($$0, 1, RLENGTH - 1); \
				printf "%s: %s\n", cmd, help_msg; \
			} \
		} \
	' $(MAKEFILE_LIST)

## Build up CLI tool.
build/cli:
	go build -o ${GOBIN}/${PROJECTNAME}-cli ${CLI_TOOL}

## Run tests.
test:
	go test ./...

## Run integration tests
test/integration:
	go test -tags=integration ./...

## Get test coverage.
test/coverage:
	go test -coverprofile=coverage.out ./...

## Download missing go dependencies.
deps:
	go mod download

## Tidy up dependencies.
tidy:
	go mod tidy

## Run go vet.
vet:
	go vet ./...

## Run linter.
lint:
	golangci-lint run ./...
