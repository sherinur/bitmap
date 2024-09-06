package errors

// errors is a package for working with errors.

import "errors"

var (
	ErrNotBMPFile      = errors.New("Error: The file is not a valid BMP file")
	ErrUnsupportedBits = errors.New("Error: Unsupported bits per pixel, only 24-bit BMP is supported")
	ErrCompressedBMP   = errors.New("Error: Compressed BMP files are not supported")
)
