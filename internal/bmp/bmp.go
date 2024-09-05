package bmp

import "fmt"

// bmp.go stores the BMP image data

// BMPFile stores the information of a bitmap file
type BMPFile struct {
	Header     BMPHeader
	InfoHeader DIBHeader
	PixelData  []byte
}

// BMPHeader stores the header of a file
type BMPHeader struct {
	Type      [2]byte
	Size      uint32
	Reserved1 uint16
	Reserved2 uint16
	Offset    uint32
}

// DIBHeader stores the dib header of a file
type DIBHeader struct {
	Size            uint32
	Width           int32
	Height          int32
	Planes          uint16
	BitsPerPixel    uint16
	Compression     uint32
	ImageSize       uint32
	XPixelsPerMeter int32
	YPixelsPerMeter int32
	ColorsUsed      uint32
	ColorsImportant uint32
}

// bmp.PrintHeader prints the extracted BMP file header.
func PrintHeader(bmpFile *BMPFile) {
	textToPrint := fmt.Sprintf(`BMP Header:
- FileType: %c%c
- FileSizeInBytes %d
- Reserved1: %d
- Reserved2: %d
- HeaderSize %d
DIB Header:
- DibHeaderSize: %d
- WidthInPixels: %d
- HeightInPixels: %d
- Planes: %d
- PixelSizeInBits: %d
- Compression: %d
- ImageSizeInBytes: %d
- XPixelsPerMeter: %d
- YPixelsPerMeter: %d
- ColorsUsed: %d
- ColorsImportant: %d`,
		bmpFile.Header.Type[0], bmpFile.Header.Type[1],
		bmpFile.Header.Size,
		bmpFile.Header.Reserved1,
		bmpFile.Header.Reserved2,
		bmpFile.Header.Offset,
		bmpFile.InfoHeader.Size,
		bmpFile.InfoHeader.Width,
		bmpFile.InfoHeader.Height,
		bmpFile.InfoHeader.Planes,
		bmpFile.InfoHeader.BitsPerPixel,
		bmpFile.InfoHeader.Compression,
		bmpFile.InfoHeader.ImageSize,
		bmpFile.InfoHeader.XPixelsPerMeter,
		bmpFile.InfoHeader.YPixelsPerMeter,
		bmpFile.InfoHeader.ColorsUsed,
		bmpFile.InfoHeader.ColorsImportant)

	fmt.Println(textToPrint)
}
