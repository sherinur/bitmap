package mirror

import "bitmap/pkg/bmp"

func Execute(inFile, outFile string) error {
	p := bmp.BitmapParser{}
	bmpFile, err := p.Parse(inFile)
	if err != nil {
		return err
	}

	bmpFile.DebugPrint()
	// err = bmpFile.ApplyMirrorVertical()
	// if err != nil {
	// 	return err
	// }

	return nil
}
