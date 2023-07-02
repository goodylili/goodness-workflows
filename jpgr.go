package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

// ToJpeg converts a PNG image to JPEG format
func ToJpeg(imageBytes []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func main() {
	directoryPath := "/Users/chukwuemeriwoukeje/Downloads" // Replace "<username>" with your actual username

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		extension := filepath.Ext(path)
		if extension == ".png" {
			imageBytes, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Failed to read image file: %s", err)
				return nil
			}

			jpegBytes, err := ToJpeg(imageBytes)
			if err != nil {
				log.Printf("Failed to convert image: %s", err)
				return nil
			}

			baseName := filepath.Base(path)
			outputPath := filepath.Join(directoryPath, baseName[:len(baseName)-len(extension)]+".jpg")

			err = os.WriteFile(outputPath, jpegBytes, os.ModePerm)
			if err != nil {
				log.Printf("Failed to write JPEG file: %s", err)
				return nil
			}

			err = os.RemoveAll(path)
			if err != nil {
				log.Printf("Failed to delete PNG file: %s", err)
				return nil
			}

			fmt.Printf("Image conversion successful: %s\n", outputPath)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error processing directory: %s", err)
	}
}
