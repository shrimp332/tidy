BIN = tidy
SRC = $(shell find . -name '*.go')

all: build

build: bin/$(BIN)

clean:
	@rm ./bin/$(BIN)

test:
	@go test -v ./test/...

install: build
	@cp ./bin/$(BIN) /usr/local/bin

install-local: build
	@mkdir -p ~/.local/bin
	@cp ./bin/$(BIN) ~/.local/bin

uninstall: /usr/local/bin/$(BIN)
	@rm /usr/local/bin/$(BIN)

uninstall-local: ~/.local/bin/$(BIN)
	@rm ~/.local/bin/$(BIN)

bin/$(BIN): $(SRC)
	@mkdir -p ./bin
	@go mod tidy
	@go build -o ./bin/$(BIN) ./cmd/$(BIN)/main.go

.PHONY: all build run test clean install install-local uninstall-local
