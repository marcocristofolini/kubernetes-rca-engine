SHELL := /bin/sh

.PHONY: fmt test vet build run-example

fmt:
	gofmt -w .

test:
	go test ./...

vet:
	go vet ./...

build:
	go build -o bin/krca ./cmd/krca

run-example:
	go run ./cmd/krca analyze examples/evicted-ephemeral-storage.json
