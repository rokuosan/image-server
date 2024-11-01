package imageprocessor

import "fmt"

func ParseSizeNxM(size string) (int, int, error) {
	var w, h int
	if _, err := fmt.Sscanf(size, "%dx%d", &w, &h); err != nil {
		return 0, 0, err
	}
	return w, h, nil
}
