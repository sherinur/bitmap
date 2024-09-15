package header

import (
	"bitmap/pkg/bmp"
)

// header.Execute executes the header command
func Execute(filepath string) error {
	p := bmp.BitmapParser{}

	bmpFile, err := p.Parse(filepath)
	if err != nil {
		return err
	}

	// bmpFile.DebugPrint()

	bmp.PrintHeader(bmpFile)

	return nil
}
