SHELL := /bin/bash


build:
	@go build -o bin/storgo ./cmd/storgo

run: build
	@./bin/storgo

test:
	@go test -v ./... 
