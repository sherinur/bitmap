package mirror

import "bitmap/internal/bmp"

func Execute(inFile, outFile string) error {
	header, err := bmp.ExtractHeader(inFile)
	if err != nil {
		return err
	}

	err = bmp.MirrorBMP(header, inFile, outFile)
	if err != nil {
		return err
	}

	return nil
}
