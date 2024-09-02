package bmp

import (
	"encoding/binary"
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

// bmp.ExtractHeader opens the file, reads and returns extracted BMP file header.
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
		return nil, fmt.Errorf("Error: %s is not bitmap file", filepath)
	}

	return &header, nil
}

// isBM checks if the file type is BM and returns boolean.
func isBM(FileType uint16) bool {
	if FileType == 0x4d42 || FileType == 0x4362 {
		return true
	}

	return false
}

// bmp.PrintHeader prints the extracted BMP file header.
func PrintHeader(h *BMPHeader) {
	textToPrint := fmt.Sprintf(`BMP Header:
- FileType BM
- FileSizeInBytes %d
- HeaderSize %d
DIB Header:
- DibHeaderSize %d
- WidthInPixels %d
- HeightInPixels %d
- PixelSizeInBits %d
- ImageSizeInBytes %d
- XPixelsPerMeter %d
- YPixelsPerMeter %d
- ColorsUsed %d
- ColorsImportant %d`,
		h.FileSize,
		h.HeaderSize,
		h.HeaderSize,
		h.Width,
		h.Height,
		h.BitsPerPixel*8,
		h.ImageSize,
		h.XPixelsPerMeter,
		h.YPixelsPerMeter,
		h.ColorsUsed,
		h.ColorsImportant)

	fmt.Println(textToPrint)
}

func MirrorBMP(dh *BMPHeader, inFile string, outFile string) error {
	f, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer f.Close()

	// Allocate memory for the mirrored image
	mirroredImage := make([]byte, dh.ImageSize)

	// Loop through each row of the image
	for y := 0; y < int(dh.Height); y++ {
		// Loop through each pixel in the row
		for x := 0; x < int(dh.Width); x++ {
			// Calculate the pixel offset in the original image
			offset := (y*int(dh.Width) + x) * int(dh.BitsPerPixel) / 8

			// Seek to the offset in the file
			_, err := f.Seek(int64(offset), 0)
			if err != nil {
				return err
			}

			// Read from the file into the mirrored image
			buf := mirroredImage[offset : offset+int(dh.BitsPerPixel)/8]
			n, err := f.Read(buf)
			if err != nil {
				return err
			}
			if n != int(dh.BitsPerPixel)/8 {
				return fmt.Errorf("short read at offset %d", offset)
			}
		}
	}

	// Open the output file
	of, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer of.Close()

	// Write the mirrored image to the output file
	_, err = of.Write(mirroredImage)
	if err != nil {
		return err
	}

	return nil
}
