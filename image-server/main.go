package main

import (
	"errors"
	"image"
	"net/http"

	"github.com/disintegration/imaging"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", nil)
}

func ResizeImage(img image.Image, width int, height int) (image.Image, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("Invalid width or height")
	}
	m := imaging.Resize(img, width, height, imaging.Lanczos)
	return m, nil
}
