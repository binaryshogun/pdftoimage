package converter

import (
	"errors"
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
)

// Convert converts PDF file to the image files of specific format.
func (c *Converter) Convert(pdfPath, outPath, format string) (err error) {
	// Create output directory if it does not exist
	if _, err = os.Stat(outPath); os.IsNotExist(err) {
		err = os.Mkdir(outPath, 0o755)
		if err != nil {
			return fmt.Errorf("could not create output directory: %v", err)
		}
	}
	defer func() {
		rmErr := os.RemoveAll(outPath)
		if rmErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to remove out directory [path=%s]: %w", outPath, rmErr))
		}
	}()

	// Scan PDF file to images.
	var images []image.Image
	images, err = c.scanner.Scan(pdfPath)
	if err != nil {
		return
	}

	// Iterate over images and save each PDF page as image.
	for imageNumber := range images {
		var (
			page = imageNumber + 1
			path = filepath.Join(outPath, fmt.Sprintf("page_%d.%s", page, format))
		)

		err = c.writer.Write(images[imageNumber], path, format)
		if err != nil {
			return
		}

		log.Printf("Page %d converted to image successfully!\n", page)
	}

	return nil
}
