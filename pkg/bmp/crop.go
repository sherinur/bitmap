package bmp

// cropImage() performs the actual cropping of the BMP image
func (bmpFile *BMPFile) CropImage(OffsetX, OffsetY, Width, Height int) {
	croppedData := make([][]Pixel, Height)
	for i := 0; i < Height; i++ {
		croppedData[i] = make([]Pixel, Width)
		copy(croppedData[i], bmpFile.ImageData[int(bmpFile.InfoHeader.Height)-Height-OffsetY+i][OffsetX:OffsetX+Width])
	}
	// update image data and DIB header
	bmpFile.ImageData = croppedData
	bmpFile.InfoHeader.Width = int32(Width)
	bmpFile.InfoHeader.Height = int32(Height)
	bmpFile.InfoHeader.ImageSize = uint32(Width * Height * int(bmpFile.InfoHeader.BitsPerPixel/8))

	// TODO:  update file size
	// bmpFile.Header.Size =
}
