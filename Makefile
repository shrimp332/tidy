BIN = tidy
SRC = $(shell find . -name '*.go')
PREFIX ?= /usr/local

all: build

build: bin/$(BIN)
	@mkdir -p ./completions/bash
	@mkdir -p ./completions/zsh
	@./bin/$(BIN) completion bash > completions/bash/tidy
	@./bin/$(BIN) completion zsh > completions/zsh/_tidy

clean:
	@rm ./bin/$(BIN)

test:
	@go test -v ./...

install: build
	@cp ./bin/$(BIN) $(PREFIX)/bin

uninstall: $(PREFIX)/bin/$(BIN)
	@rm $(PREFIX)/bin/$(BIN)

bin/$(BIN): $(SRC)
	@mkdir -p ./bin
	@go mod tidy
	@go build -o ./bin/$(BIN) .

.PHONY: all build test run clean install
