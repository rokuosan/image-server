package main

import (
	"errors"
	"fmt"
	"image"

	"github.com/disintegration/imaging"
)

func main() {
	fmt.Println("Hello, World!")
}

func ResizeImage(img image.Image, width int, height int) (image.Image, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("Invalid width or height")
	}
	m := imaging.Resize(img, width, height, imaging.Lanczos)
	return m, nil
}
