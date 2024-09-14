SHELL := /bin/bash


build:
	@go build -o bin/storgo ./cmd/storgo

run: build
	@./bin/fs

test:
	@go test -v ./... 