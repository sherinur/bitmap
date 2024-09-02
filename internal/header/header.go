package header

import (
	"bitmap/internal/bmp"
)

func Execute(filepath string) error {
	header, err := bmp.ExtractHeader(filepath)
	if err != nil {
		return err
	}

	bmp.PrintHeader(header)

	return nil
}
