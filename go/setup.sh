#!/bin/bash

# Setup script for image Rainfall Go Application

set -e

echo "🚀 Setting up image Rainfall Go Application..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    echo "Visit: https://golang.org/doc/install"
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | grep -o 'go[0-9]\+\.[0-9]\+' | sed 's/go//')
REQUIRED_VERSION="1.21"

if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]; then
    echo "❌ Go version $GO_VERSION is too old. Please upgrade to Go 1.21 or higher."
    exit 1
fi

echo "✅ Go version $GO_VERSION detected"

# Download dependencies
echo "📦 Installing dependencies..."
go mod tidy

# Build the application
echo "🔨 Building application..."
go build -o image-server main.go

echo "✅ Setup complete!"
echo ""
echo "To run the application:"
echo "  ./image-server"
echo ""
echo "Or use make commands:"
echo "  make run      # Run development server"
echo "  make convert  # Convert images to WebP"
echo "  make build    # Build production binary"
