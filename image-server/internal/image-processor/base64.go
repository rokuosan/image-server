package imageprocessor

import (
	"encoding/base64"
	"image"
	"strings"
)

func Base64ToImage(message string) (image.Image, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(message))
	m, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return m, nil
}
