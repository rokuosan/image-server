package imageprocessor

import (
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
