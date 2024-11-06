BIN = tidy
SRC = $(shell find . -name '*.go')

all: build

build: $(BIN)

clean:
	@rm $(BIN)

test:
	@go test -v ./test/...

install: build
	@cp ./$(BIN) /usr/local/bin

install-local: build
	@mkdir -p ~/.local/bin
	@cp ./$(BIN) ~/.local/bin

uninstall: /usr/local/bin/$(BIN)
	@rm /usr/local/bin/$(BIN)

uninstall-local: ~/.local/bin/$(BIN)
	@rm ~/.local/bin/$(BIN)

$(BIN): $(SRC)
	@go mod tidy
	@go build -o $(BIN) ./cmd/main.go

.PHONY: all build run test clean install install-local uninstall-local
