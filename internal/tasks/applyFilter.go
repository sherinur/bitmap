package tasks

import (
	"fmt"

	"bitmap/pkg/bmp"
)

// execute function (input file, file that we need to create, color to apply)
func ApplyFilter(bmpFile *bmp.BMPFile, value string) error {
	switch value {
	case "red", "green", "blue":
		bmpFile.ApplyByColor(value)

	case "grayscale":
		bmpFile.ApplyGrayScale()

	case "negative":
		bmpFile.ApplyNegative()

	case "pixelate":
		bmpFile.ApplyPixelation(20)

	case "blur":
		bmpFile.ApplyBlur()

	default:
		return fmt.Errorf("Filter Error: unknown value or filter: %s", value)
	}

	return nil
}
