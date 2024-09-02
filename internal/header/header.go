package header

import (
	"bitmap/internal/bmp"
)

// header.Execute executes the header command
func Execute(filepath string) error {
	header, err := bmp.ExtractHeader(filepath)
	if err != nil {
		return err
	}

	bmp.PrintHeader(header)

	return nil
}
