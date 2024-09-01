package header

import (
	"fmt"

	"bitmap/internal/bmp"
)

func Execute(filepath string) error {
	header, err := bmp.ExtractHeader(filepath)
	if err != nil {
		return err
	}

	fmt.Println(header)

	return nil
}
