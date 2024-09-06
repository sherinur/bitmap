package bmp

func (bmpFile *BMPFile) ApplyMirrorVertical() error {
	rowsNum := len(bmpFile.ImageData)
	for i := 0; i < rowsNum; i++ {
		bmpFile.ImageData[rowsNum-1], bmpFile.ImageData[i] = bmpFile.ImageData[i], bmpFile.ImageData[rowsNum-1]
	}

	bmpFile.ImageData[0][10].Red = 0
	return nil
}
