package scanner

import (
	"fmt"
	"os"
	"testing"

	"github.com/gen2brain/go-fitz"
	"github.com/jung-kurt/gofpdf"
	"github.com/stretchr/testify/assert"
)

func TestScanner_Scan(t *testing.T) {
	const path = "./test.pdf"

	tests := []struct {
		name  string
		pages int
		err   error
	}{
		{
			name:  "success_one_page",
			pages: 1,
			err:   nil,
		},
		{
			name:  "success_two_pages",
			pages: 2,
			err:   nil,
		},
		{
			name:  "success_five_pages",
			pages: 5,
			err:   nil,
		},
		{
			name:  "error_no_document",
			pages: 0,
			err:   fmt.Errorf("could not open PDF file: %w", fitz.ErrNoSuchFile),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pages != 0 {
				pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitMillimeter, gofpdf.PageSizeA4, "")

				for i := 0; i < tt.pages; i++ {
					pdf.AddPage()
				}

				err := pdf.OutputFileAndClose(path)
				assert.NoError(t, err)

				t.Cleanup(func() {
					err = os.RemoveAll(path)
					assert.NoError(t, err)
				})
			}

			scanner := NewScanner()

			images, err := scanner.Scan(path)
			assert.Equal(t, tt.pages, len(images))
			assert.Equal(t, tt.err, err)
		})
	}
}
