package bmp

func NewBMPFile() *BMPFile {
	return &BMPFile{
		Header:     BMPHeader{},
		InfoHeader: DIBHeader{},
		ImageData:  make([][]Pixel, 0),
	}
}
