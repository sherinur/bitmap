package tasks

import (
	"errors"
	"strconv"
	"strings"

	"bitmap/pkg/bmp"
)

func ApplyCrop(bmpFile *bmp.BMPFile, cropParams string) error {
	// split crop parameters
	parts := strings.Split(cropParams, "-")
	var OffsetX, OffsetY, Width, Height int
	var err error

	// validate and parse crop parameters~
	if len(parts) == 2 || len(parts) == 4 {
		OffsetX, err = strconv.Atoi(parts[0])
		if err != nil {
			return ErrInvalidOffsetX
		}
		OffsetY, err = strconv.Atoi(parts[1])
		if err != nil {
			return ErrInvalidOffsetY
		}

		// default width and height if not provided
		if len(parts) == 2 {
			Width = int(bmpFile.InfoHeader.Width) - OffsetX
			Height = int(bmpFile.InfoHeader.Height) - OffsetY
		} else {
			Width, err = strconv.Atoi(parts[2])
			if err != nil {
				return ErrInvalidWidth
			}
			Height, err = strconv.Atoi(parts[3])
			if err != nil {
				return ErrInvalidHeight
			}
		}
	} else {
		return ErrInvalidFormat
	}

	// check if crop area is within image bounds
	if OffsetX < 0 || OffsetY < 0 || OffsetX+Width > int(bmpFile.InfoHeader.Width) || OffsetY+Height > int(bmpFile.InfoHeader.Height) || OffsetX > int(bmpFile.InfoHeader.Width) || OffsetY > int(bmpFile.InfoHeader.Height) {
		return errors.New("crop area exceeds image boundaries")
	}

	// perform cropping
	bmpFile.CropImage(OffsetX, OffsetY, Width, Height)

	return nil
}
