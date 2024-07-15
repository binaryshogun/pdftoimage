package writer

import (
	"fmt"
	"image"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriter_Write(t *testing.T) {
	const path = "./image"

	tests := []struct {
		name      string
		format    string
		checkFile bool
		err       error
	}{
		{
			name:      "success_png_image",
			format:    png,
			checkFile: true,
			err:       nil,
		},
		{
			name:      "success_jpeg_image",
			format:    jpeg,
			checkFile: true,
			err:       nil,
		},
		{
			name:      "success_jpg_image",
			format:    jpg,
			checkFile: true,
			err:       nil,
		},
		{
			name:      "error_wrong_format",
			format:    "gif",
			checkFile: false,
			err:       errorUnknownFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				writer    = NewWriter()
				imagePath = fmt.Sprintf("%s.%s", path, tt.format)
				img       = image.NewNRGBA(image.Rect(0, 0, 100, 100))
			)

			t.Cleanup(func() {
				err := os.RemoveAll(imagePath)
				assert.NoError(t, err)
			})

			err := writer.Write(img, imagePath, tt.format)
			assert.Equal(t, tt.err, err, "Errors should be equal")

			if tt.checkFile {
				_, err = os.Stat(imagePath)
				assert.NoError(t, err)
			}
		})
	}
}
