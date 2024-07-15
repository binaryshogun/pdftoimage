package converter

import (
	"fmt"
	"image"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverter_Convert(t *testing.T) {
	const (
		pdfPath    = "./test.pdf"
		outPath    = "images"
		formatJPEG = "jpeg"
		formatPNG  = "png"
	)

	tests := []struct {
		name    string
		pdfPath string
		outPath string
		format  string
		setup   func(
			psMock *pdfScannerMock,
			iwMock *imageWriterMock,
			pdfPath, outPath, format string,
		)
		wantErr error
	}{
		{
			name:    "convert_pdf_and_write_jpeg_images_no_error",
			pdfPath: pdfPath,
			outPath: outPath,
			format:  formatJPEG,
			setup: func(
				psMock *pdfScannerMock,
				iwMock *imageWriterMock,
				pdfPath, outPath, format string,
			) {
				images := []image.Image{
					image.NewRGBA(image.Rect(0, 0, 1, 1)),
					image.NewRGBA(image.Rect(0, 0, 1, 1)),
				}

				psMock.On("Scan", pdfPath).Return(images, nil).Once()

				iwMock.On("Write", images[0], fmt.Sprintf("%s/%s.%s", outPath, "page_1", format), format).Return(nil).Once()
				iwMock.On("Write", images[1], fmt.Sprintf("%s/%s.%s", outPath, "page_2", format), format).Return(nil).Once()
			},
			wantErr: nil,
		},
		// TODO: add more test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				ps = newPdfScannerMock(t)
				iw = newImageWriterMock(t)
			)

			tt.setup(ps, iw, tt.pdfPath, tt.outPath, tt.format)

			converter := NewConverter(ps, iw)

			err := converter.Convert(tt.pdfPath, tt.outPath, tt.format)
			defer func() {
				assert.NoError(t, os.RemoveAll(outPath))
			}()

			assert.Equal(t, tt.wantErr, err)

			ps.AssertExpectations(t)
			iw.AssertExpectations(t)
		})
	}
}
