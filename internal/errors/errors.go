package errors

// errors is a package for working with errors.

import "errors"

var (
	ErrNotBMPFile      = errors.New("the file is not a valid BMP file")
	ErrUnsupportedBits = errors.New("unsupported bits per pixel, only 24-bit BMP is supported")
	ErrCompressedBMP   = errors.New("compressed BMP files are not supported")
)
