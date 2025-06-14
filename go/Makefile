# Makefile for image Rainfall Go Application

.PHONY: help install run build convert clean test

# Default target
help:
	@echo "Available commands:"
	@echo "  install  - Download and install dependencies"
	@echo "  run      - Run the web server"
	@echo "  build    - Build the application binary"
	@echo "  convert  - Run image conversion utility"
	@echo "  clean    - Clean build artifacts"
	@echo "  test     - Run tests"

# Install dependencies
install:
	go mod tidy
	go mod download

# Run the web server
run:
	go run main.go

# Build the application
build:
	go build -o image-server main.go

# Run image conversion
convert:
	go run convert.go

# Clean build artifacts
clean:
	rm -f image-server
	go clean

# Run tests (if any)
test:
	go test ./...
