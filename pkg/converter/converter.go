package converter

import (
	"errors"
	"fmt"
	"github.com/binaryshogun/pdftoimage/pkg/pdf"
	"image"
	jpegEncoder "image/jpeg"
	pngEncoder "image/png"
	"io"
	"os"
	"path/filepath"
)

const (
	jpeg = "jpeg"
	jpg  = "jpg"
	png  = "png"
)

var (
	// errorUnknownFormat specifies wrong output format.
	errorUnknownFormat = fmt.Errorf("unknown output format")
)

type (
	// scanner - PDF files scanner.
	scanner interface {
		Scan(path string) ([]image.Image, error)
	}

	// Converter - PDF to image converter.
	Converter interface {
		Convert(pdfPath, outPath, format string) error
	}
)

type (
	converter struct {
		scanner scanner
	}

	encoder func(w io.Writer, img image.Image) error
)

// NewConverter creates new converter by output format.
func NewConverter() Converter {
	return &converter{scanner: pdf.NewScanner()}
}

func (c *converter) Convert(pdfPath, outPath, format string) (err error) {
	// Check output image format.
	var encodeFunc encoder
	encodeFunc, err = getEncodeFuncForImage(format)
	if err != nil {
		return
	}

	// Create output directory if it does not exist
	if _, err = os.Stat(outPath); os.IsNotExist(err) {
		err = os.Mkdir(outPath, 0755)
		if err != nil {
			return fmt.Errorf("could not create output directory: %v", err)
		}
	}

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

		err = saveImage(images[imageNumber], path, encodeFunc)
		if err != nil {
			return
		}

		fmt.Printf("Page %d converted to image successfully!\n", page)
	}

	return nil
}

// getEncodeFuncForImage gets the encoding function for the image.
func getEncodeFuncForImage(format string) (encoder, error) {
	switch format {
	case jpeg, jpg:
		return func(w io.Writer, img image.Image) error {
			err := jpegEncoder.Encode(w, img, &jpegEncoder.Options{Quality: 100})
			if err != nil {
				return err
			}

			return nil
		}, nil

	case png:
		return func(w io.Writer, img image.Image) error {
			err := pngEncoder.Encode(w, img)
			if err != nil {
				return err
			}

			return nil
		}, nil

	default:
		return nil, errorUnknownFormat
	}
}

// saveImage saves images on disc.
func saveImage(img image.Image, path string, encodeFunc encoder) error {
	// Create file on path.
	outFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create output file: %v", err)
	}
	defer func(outFile *os.File) {
		closeErr := outFile.Close()
		if err != nil {
			err = errors.Join(err, fmt.Errorf("failed to close image file [path=%s]: %w", path, closeErr))
		}
	}(outFile)

	// Encode image using encoding function.
	err = encodeFunc(outFile, img)
	if err != nil {
		return fmt.Errorf("could not encode image: %v", err)
	}

	return nil
}
