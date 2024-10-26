package converter

import (
	"image"

	"golang.org/x/image/draw"
)

func ResizeImage(img image.Image, width int, height int) (image.Image, error) {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.BiLinear.Scale(m, m.Bounds(), img, img.Bounds(), draw.Over, nil)
	return m, nil
}
