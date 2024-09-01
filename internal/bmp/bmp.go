package bmp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type BMPHeader struct {
	FileType        uint16
	FileSize        uint32
	Reserved1       uint16
	Reserved2       uint16
	Offset          uint32
	HeaderSize      uint32
	Width           int32
	Height          int32
	Planes          uint16
	BitsPerPixel    uint16
	Compression     uint32
	ImageSize       uint32
	XPixelsPerMeter uint32
	YPixelsPerMeter uint32
	ColorsUsed      uint32
	ColorsImportant uint32
}

func ExtractHeader(filepath string) (*BMPHeader, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var header BMPHeader
	err = binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		return nil, err
	}

	if !isBM(header.FileType) {
		return nil, errors.New("File format not matching BM")
	}

	return &header, nil
}

func isBM(FileType uint16) bool {
	if FileType == 0x4d42 || FileType == 0x4362 {
		return true
	}

	return false
}

func (h *BMPHeader) String() string {
	return fmt.Sprintf(`BMP Header:
- FileType BM
- FileSizeInBytes %d
- HeaderSize %d
DIB Header:
- DibHeaderSize %d
- WidthInPixels %d
- HeightInPixels %d
- PixelSizeInBits %d
- ImageSizeInBytes %d`,
		h.FileSize,
		h.HeaderSize,
		h.HeaderSize,
		h.Width,
		h.Height,
		h.BitsPerPixel*8,
		h.ImageSize)
}
