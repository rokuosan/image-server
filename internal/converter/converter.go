package converter

import (
	"bytes"
	"errors"
	"image"

	"github.com/disintegration/imaging"
)

func ResizeImage(img image.Image, width int, height int) (image.Image, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("Invalid width or height")
	}
	m := imaging.Resize(img, width, height, imaging.Lanczos)
	return m, nil
}

func DeleteExif(img image.Image, format imaging.Format) (image.Image, error) {
	r := new(bytes.Buffer)
	if err := imaging.Encode(r, img, format); err != nil {
		return nil, err
	}

	m, err := imaging.Decode(r, imaging.AutoOrientation(true))
	if err != nil {
		return nil, err
	}

	return m, nil
}
