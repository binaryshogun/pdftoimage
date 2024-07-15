//go:build integration

package converter

import (
	"testing"

	"github.com/binaryshogun/pdftoimage/pkg/pdf/converter"
	"github.com/stretchr/testify/assert"
)

func TestConverter_Convert(t *testing.T) {
	// TODO: check convert function on the real PDF file and assert there are images.

	c := converter.Converter{}

	assert.Error(t, c.Convert("", "", ""))
}
