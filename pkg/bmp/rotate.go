package bmp

import (
	"fmt"
)

// ApplyRotation calculates the final rotation angle and applies the corresponding rotation to the BMP file.
func (bmpFile *BMPFile) ApplyRotation(rotation string) error {
	rotationAngles := map[string]int{
		"right": -90,
		"left":  90,
		"90":    -90,
		"-90":   90,
		"180":   180,
		"-180":  -180,
		"270":   -270,
		"-270":  270,
	}

	angle := 0
	if delta, ok := rotationAngles[rotation]; ok {
		angle = (angle + delta) % 360
	} else {
		return fmt.Errorf("invalid rotation option: %s", rotation)
	}

	// Normalize the angle to be non-negative
	if angle < 0 {
		angle += 360
	}

	switch angle {
	case 0:
		// No rotation needed
	case 90:
		bmpFile.ImageData = rotate90(bmpFile)
	case 180:
		bmpFile.ImageData = rotate180(bmpFile)
	case 270:
		bmpFile.ImageData = rotate270(bmpFile)
	default:
		return fmt.Errorf("unsupported rotation angle: %d", angle)
	}

	return nil
}

// rotate90 rotates the image data 90 degrees clockwise.
func rotate90(bmpFile *BMPFile) [][]Pixel {
	oldWidth := int(bmpFile.InfoHeader.Width)
	oldHeight := int(bmpFile.InfoHeader.Height)

	// Allocate new image data with swapped width and height
	newImageData := make([][]Pixel, oldWidth)
	for i := range newImageData {
		newImageData[i] = make([]Pixel, oldHeight)
	}

	// Transpose the image data to achieve 90-degree rotation
	for y := 0; y < oldHeight; y++ {
		for x := 0; x < oldWidth; x++ {
			newX := oldHeight - 1 - y
			newY := x
			newImageData[newY][newX] = bmpFile.ImageData[y][x]
		}
	}

	// Update BMP dimensions after rotation
	bmpFile.InfoHeader.Width = int32(oldHeight)
	bmpFile.InfoHeader.Height = int32(oldWidth)
	return newImageData
}

// rotate180 rotates the image data 180 degrees using two 90-degree rotations.
func rotate180(bmpFile *BMPFile) [][]Pixel {
	for i := 0; i < 2; i++ {
		bmpFile.ImageData = rotate90(bmpFile)
	}
	return bmpFile.ImageData
}

// rotate270 rotates the image data 270 degrees clockwise using three 90-degree rotations.
func rotate270(bmpFile *BMPFile) [][]Pixel {
	for i := 0; i < 3; i++ {
		bmpFile.ImageData = rotate90(bmpFile)
	}
	return bmpFile.ImageData
}
