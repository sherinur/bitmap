package header

import (
	"fmt"
	"os"

	"bitmap/internal/bmp"
)

// header.Execute executes the header command
func Execute(filepath string) error {
	p := bmp.BitmapParser{}

	bmpFile, err := p.Parse(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bmp.PrintHeader(bmpFile)

	return nil
}
