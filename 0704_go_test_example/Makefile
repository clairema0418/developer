.PHONY: build run test

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
BINARY_NAME=main

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

run: 
	$(GORUN) main.go

test: 
	$(GOTEST) -v ./...