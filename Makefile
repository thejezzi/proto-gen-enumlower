.PHONY: all build test generate clean

all: build test

build:
	go build ./...

test:
	go test ./...

clean:
	rm -rf bin
