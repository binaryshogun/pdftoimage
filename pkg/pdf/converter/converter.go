package converter

import (
	"image"
)

//go:generate mockery --name pdfScanner --output . --outpkg converter --filename pdf_scanner_mock.go --structname pdfScannerMock
//go:generate mockery --name imageWriter --output . --outpkg converter --filename image_writer_mock.go --structname imageWriterMock

type (
	pdfScanner interface {
		Scan(path string) ([]image.Image, error)
	}
	imageWriter interface {
		Write(img image.Image, path, format string) error
	}
)

// Converter - PDF to image converter.
type Converter struct {
	scanner pdfScanner
	writer  imageWriter
}

// NewConverter creates new converter by output format.
func NewConverter(scanner pdfScanner, writer imageWriter) *Converter {
	return &Converter{scanner: scanner, writer: writer}
}
