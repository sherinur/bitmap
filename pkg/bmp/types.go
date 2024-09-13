package bmp

// bmp.go stores the BMP image data

// BMPFile stores the information of a bitmap file
type BMPFile struct {
	Header     BMPHeader
	InfoHeader DIBHeader
	ImageData  [][]Pixel
}

type BMPManipulator interface {
	ApplyNegative()

	ApplyGrayScale()

	ApplyByColor(string)

	ApplyBlur()

	ApplyPixelation(int)

	CropImage(int, int, int, int)

	ApplyMirrorHorizontal()

	ApplyMirrorVertical()

	ApplyRotation(string)

	DebugPrint()
}

// Pixel is used for BMPFile.Pixels
type Pixel struct {
	Blue  uint8
	Green uint8
	Red   uint8
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
