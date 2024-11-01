package imageprocessor

import (
	"bytes"
	"errors"
	"image"

	"github.com/disintegration/imaging"
)

func EncodeImage(img image.Image, format string) ([]byte, error) {
	if format != "jpeg" && format != "png" {
		return nil, errors.New("Unsupported format: " + format)
	}
	buf := new(bytes.Buffer)
	f, err := imaging.FormatFromExtension(format)
	if err != nil {
		return nil, err
	}
	imaging.Encode(buf, img, f)

	return buf.Bytes(), nil
}
