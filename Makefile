BIN = tidy
SRC = $(shell find . -name '*.go')
PREFIX ?= /usr/local

all: build

build: bin/$(BIN)

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
