package writer

import (
	"errors"
	"fmt"
	"image"
	jpegEncoder "image/jpeg"
	pngEncoder "image/png"
	"io"
	"os"
)

type encoder func(w io.Writer, img image.Image) error

// Write saves images on disc.
func (w *Writer) Write(img image.Image, path, format string) error {
	// Check output image format.
	encodeFunc, err := getEncodeFuncForImage(format)
	if err != nil {
		return err
	}

	// Create file on path.
	var outFile *os.File
	outFile, err = os.Create(path)
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
