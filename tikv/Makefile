-include .env

PROJECT_NAME := $(shell basename "$(PWD)")
PROJECT_DIR := $(shell pwd)
BUILD_DIR := ./cmd/tikv/server

# Basic go command
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary name
BINARY_NAME=tikv
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(BUILD_DIR)/main.go
run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(BUILD_DIR)/*.go
	./$(BINARY_NAME) -grpc-port=$(PORT)
test:
	$(GOTEST) -v $(PROJECT_DIR)/test
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)
	rm -rf $(BINARY_UNIX)
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v