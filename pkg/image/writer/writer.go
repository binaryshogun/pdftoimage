package writer

import "fmt"

const (
	jpeg = "jpeg"
	jpg  = "jpg"
	png  = "png"
)

// errorUnknownFormat specifies wrong output format.
var errorUnknownFormat = fmt.Errorf("unknown output format")

type Writer struct{}

func NewWriter() *Writer {
	return &Writer{}
}
