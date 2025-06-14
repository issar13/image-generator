package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
	"golang.org/x/image/draw"
)

const (
	inputDir    = "../public/images"
	outputDir   = "../public/images/compressed-webp"
	maxSize     = 100 * 1024 // 100KB
	maxWidth    = 1280
	maxHeight   = 1280
	webpQuality = 60
)

// isImageFile checks if the file is a supported image format
func isImageFile(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

// getFileSize returns the size of a file in bytes
func getFileSize(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// loadImage loads an image from file
func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".jpg", ".jpeg":
		return jpeg.Decode(file)
	case ".png":
		return png.Decode(file)
	default:
		img, _, err := image.Decode(file)
		return img, err
	}
}

// resizeImage resizes an image while maintaining aspect ratio
func resizeImage(img image.Image, maxWidth, maxHeight int) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Calculate new dimensions while maintaining aspect ratio
	if width <= maxWidth && height <= maxHeight {
		return img
	}

	ratio := float64(width) / float64(height)
	newWidth := maxWidth
	newHeight := int(float64(newWidth) / ratio)

	if newHeight > maxHeight {
		newHeight = maxHeight
		newWidth = int(float64(newHeight) * ratio)
	}

	// Create new image with calculated dimensions
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, bounds, draw.Over, nil)

	return dst
}

// compressToWebp converts and compresses an image to WebP format
func compressToWebp(inputFile, outputFile string) error {
	// Check if file is already small enough
	size, err := getFileSize(inputFile)
	if err != nil {
		return err
	}

	if size < maxSize {
		fmt.Printf("⚠️ Skipped (already small): %s\n", filepath.Base(inputFile))
		return nil
	}

	// Load the image
	img, err := loadImage(inputFile)
	if err != nil {
		return fmt.Errorf("failed to load image: %v", err)
	}

	// Resize the image if necessary
	resizedImg := resizeImage(img, maxWidth, maxHeight)

	// Create output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	// Encode to WebP
	options := &webp.Options{
		Lossless: false,
		Quality:  webpQuality,
	}

	err = webp.Encode(outFile, resizedImg, options)
	if err != nil {
		return fmt.Errorf("failed to encode WebP: %v", err)
	}

	// Get new file size
	newSize, err := getFileSize(outputFile)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Compressed: %s → %d KB\n", filepath.Base(inputFile), newSize/1024)
	return nil
}

func main() {
	fmt.Printf("📁 Scanning folder: %s\n", filepath.Clean(inputDir))

	// Ensure output directory exists
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Read input directory
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatalf("Failed to read input directory: %v", err)
	}

	// Filter image files
	var imageFiles []string
	for _, file := range files {
		if !file.IsDir() && isImageFile(file.Name()) {
			imageFiles = append(imageFiles, file.Name())
		}
	}

	fmt.Printf("🖼️ Found %d image(s) to compress.\n", len(imageFiles))

	if len(imageFiles) == 0 {
		fmt.Println("⚠️ No image files found.")
		return
	}

	// Process each image file
	for _, file := range imageFiles {
		inputFile := filepath.Join(inputDir, file)
		baseName := strings.TrimSuffix(file, filepath.Ext(file))
		outputFile := filepath.Join(outputDir, baseName+".webp")

		err := compressToWebp(inputFile, outputFile)
		if err != nil {
			fmt.Printf("❌ Failed to compress %s: %v\n", file, err)
		}
	}

	fmt.Println("🎉 Compression completed!")
}
