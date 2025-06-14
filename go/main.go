package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

const (
	PORT      = 3000
	IMAGE_DIR = "../public/images"
)

// ImageResponse represents the JSON response for images
type ImageResponse struct {
	Images []string `json:"images,omitempty"`
	Error  string   `json:"error,omitempty"`
}

// getImageFiles returns a list of image files from the images directory
func getImageFiles(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Read the images directory
	files, err := os.ReadDir(IMAGE_DIR)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := ImageResponse{Error: "Unable to read image directory"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Filter image files using regex
	imageRegex := regexp.MustCompile(`(?i)\.(jpe?g|png|gif|webp|bmp|heic)$`)
	var imageFiles []string

	for _, file := range files {
		if !file.IsDir() && imageRegex.MatchString(file.Name()) {
			imageFiles = append(imageFiles, "images/"+file.Name())
		}
	}

	// Return the list of image files
	json.NewEncoder(w).Encode(imageFiles)
}

// serveStaticFiles serves static files from the public directory
func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	// Remove leading slash from URL path
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		path = "index.html"
	}

	// Construct the full file path
	fullPath := filepath.Join("../public", path)

	// Check if file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	// Determine content type based on file extension
	ext := filepath.Ext(fullPath)
	switch ext {
	case ".html":
		w.Header().Set("Content-Type", "text/html")
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".webp":
		w.Header().Set("Content-Type", "image/webp")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
	}

	// Serve the file
	http.ServeFile(w, r, fullPath)
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/api/images", getImageFiles).Methods("GET")

	// Static file serving - catch all other routes
	r.PathPrefix("/").HandlerFunc(serveStaticFiles)

	// Start the server
	fmt.Printf("Server running at http://localhost:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), r))
}
