$(VERBOSE).SILENT:
SHELL := /bin/bash
CYAN := \033[0;36m
NC := \033[0m

.DEFAULT_GOAL := help

APP_NAME := gopassgen
BUILD_DIR := bin
SRC := main.go

help:
	grep -E '^[a-zA-Z_-]+:[ \t]+.*?# .*$$' $(MAKEFILE_LIST) | sort | awk -F ':.*?# ' '{printf "  ${CYAN}%-24s${NC}\t%s\n", $$1, $$2}'

run: # Run the application
	go run $(SRC)

build: # Build the application
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC)

test: # Run tests
	go test -v

test-coverage: # Run tests with coverage report
	mkdir -p $(BUILD_DIR)
	go test -coverprofile=$(BUILD_DIR)/coverage.out
	go tool cover -html=$(BUILD_DIR)/coverage.out -o $(BUILD_DIR)/coverage.html

clean: # Clean build artifacts
	rm -rf $(BUILD_DIR)

.PHONY: run build test test-coverage clean
