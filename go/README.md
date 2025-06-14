# Imraan Rainfall - Go Version

This is a Go implementation of the Imraan interactive rainfall effect web application.

## Features

- **Web Server**: Serves the interactive HTML page with rainfall effects
- **Image API**: REST API endpoint that returns a list of available images
- **Static File Serving**: Serves HTML, CSS, JS, and image files
- **Image Conversion**: Utility to compress and convert images to WebP format

## Project Structure

```text
go/
├── main.go          # Main web server application
├── convert.go       # Image compression utility
├── go.mod          # Go module dependencies
└── README.md       # This file
```

## Prerequisites

- Go 1.21 or higher
- Internet connection for downloading dependencies

## Installation

1. Navigate to the Go directory:

   ```bash
   cd go
   ```

2. Download dependencies:

   ```bash
   go mod tidy
   ```

## Running the Application

### Start the Web Server

```bash
go run main.go
```

The server will start on `http://localhost:3000`

### Convert Images (Optional)

To compress and convert images to WebP format:

```bash
go run convert.go
```

This will:

- Scan the `../public/images` directory for JPEG and PNG files
- Skip files already under 100KB
- Resize images to fit within 1280x1280 pixels
- Convert to WebP format with 60% quality
- Save compressed images to `../public/images/compressed-webp/`

## API Endpoints

- `GET /` - Serves the main HTML page
- `GET /api/images` - Returns JSON array of available image filenames
- `GET /images/*` - Serves image files
- `GET /*` - Serves static files (HTML, CSS, JS)

## Dependencies

- **gorilla/mux**: HTTP router and URL matcher
- **chai2010/webp**: WebP image encoding/decoding
- **golang.org/x/image**: Extended image processing capabilities

## Build for Production

To build a standalone executable:

```bash
go build -o imraan-server main.go
```

Then run:

```bash
./imraan-server
```

## Differences from Node.js Version

- Uses Gorilla Mux for routing instead of Express
- Native Go image processing instead of Sharp
- Better memory management and performance
- Single binary deployment (no node_modules)
- Cross-platform compilation support

## Performance Benefits

- Lower memory footprint
- Faster startup time
- Better concurrent request handling
- No runtime dependencies after compilation
