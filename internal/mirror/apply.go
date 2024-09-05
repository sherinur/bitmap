package mirror

import (
	"os"

	"bitmap/internal/bmp"
)

// bmp.MirrorBMP
func applyMirror(bmpFile *bmp.BMPFile, inFile string, outFile string) error {
	f, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer f.Close()

	// TODO: Implement mirror logic.

	return nil
}
