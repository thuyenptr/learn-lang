-include .env

PROJECT_NAME := $(shell basename "$(PWD)")
PROJECT_DIR := $(shell pwd)
BUILD_DIR := $(PROJECT_DIR)

# Basic go command
GOCMD=go
BASH=bash
GEN_PROTO_CMD=$(BASH) ${PROJECT_DIR}/scripts/gen-proto.sh
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GENPROTO=$(GEN_PROTO_CMD)

# Binary name
BINARY_NAME=zps-binlog

build:
	$(GENPROTO)
	$(GOBUILD) -o $(BINARY_NAME) -v $(BUILD_DIR)/*.go
run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(BUILD_DIR)/*.go
	./$(BINARY_NAME)
test:
	$(GOTEST) -v $(PROJECT_DIR)/test
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)
	rm ./proto/*.go
