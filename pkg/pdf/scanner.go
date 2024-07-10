package pdf

import (
	"errors"
	"fmt"
	"github.com/gen2brain/go-fitz"
	"image"
)

// Scanner - scanner for PDF documents.
type Scanner struct{}

// NewScanner creates new instance of *[Scanner].
func NewScanner() *Scanner {
	return &Scanner{}
}

// Scan scans the PDF file and converts it to the slice of images.
func (s *Scanner) Scan(path string) (images []image.Image, err error) {
	// Open the PDF file
	var doc *fitz.Document
	doc, err = fitz.New(path)
	if err != nil {
		return nil, fmt.Errorf("could not open PDF file: %w", err)
	}
	defer func(doc *fitz.Document) {
		closeErr := doc.Close()
		if closeErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to close PDF file: %w", closeErr))
		}
	}(doc)

	// Allocate space for images.
	images = make([]image.Image, doc.NumPage())

	// Iterate through each page and convert it to an image
	for pageNumber := 0; pageNumber < doc.NumPage(); pageNumber++ {
		var img image.Image
		img, err = doc.Image(pageNumber)
		if err != nil {
			return nil, fmt.Errorf("could not get image from page %d: %v", pageNumber, err)
		}

		images[pageNumber] = img
	}

	return
}
