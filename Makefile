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

uninstall: /usr/local/bin/$(BIN)
	@rm /usr/local/bin/$(BIN)

$(BIN): $(SRC)
	@go mod tidy
	@go build -o $(BIN) ./cmd/main.go

.PHONY: all build run test clean install
