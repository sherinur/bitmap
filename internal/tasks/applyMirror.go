package tasks

import (
	"strings"

	"test/pkg/bmp"
)

// Execute processes the image file, applies the mirror effect, and saves the result.
func ApplyMirror(bmpFile *bmp.BMPFile, mirrorOption string) error {
	// Determine the mirror effect based on the mirrorOption string.
	switch strings.ToLower(mirrorOption) {
	case "horizontal", "h", "horizontally", "hor":
		if err := bmpFile.ApplyMirrorHorizontal(); err != nil {
			return err
		}
	case "vertical", "v", "vertically", "ver":
		if err := bmpFile.ApplyMirrorVertical(); err != nil {
			return err
		}
	default:
		return ErrInvalidDirection
	}

	return nil
}
