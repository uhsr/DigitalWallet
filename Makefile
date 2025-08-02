# Makefile for DigitalWallet

BINARY_NAME=digitalwallet
BUILD_DIR=./build

.PHONY: build test clean run help

build: ## Build the application
	@echo "Building {BINARY_NAME}..."
	@mkdir -p {BUILD_DIR}
	go build -o {BUILD_DIR}/{BINARY_NAME} ./cmd/{BINARY_NAME}

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf {BUILD_DIR}

run: build ## Build and run the application
	@echo "Running {BINARY_NAME}..."
	@./{BUILD_DIR}/{BINARY_NAME} -verbose

help: ## Show this help message
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
