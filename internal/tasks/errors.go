package tasks

import "errors"

var (
	ErrInvalidDirection = errors.New("Mirror Error: invalid mirror direction")

	ErrInvalidOffsetY = errors.New("Crop Error: invalid crop parameters: OffsetY must be a number")
	ErrInvalidOffsetX = errors.New("Crop Error: invalid crop parameters: OffsetX must be a number")
	ErrInvalidWidth   = errors.New("Crop Error: invalid crop parameters: Width must be a number")
	ErrInvalidHeight  = errors.New("Crop Error: invalid crop parameters: Height must be a number")
	ErrInvalidFormat  = errors.New("Crop Error: invalid crop format, expected 2 or 4 values")
)
