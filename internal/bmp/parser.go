package bmp

// parser.go provides functions to parse the bitmap(.bmp, .dib) file

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"bitmap/internal/errors"
)

// BMPParser is a parser for bitmap(.bmp, .dib) file
type BMPParser interface {
	Parse(r io.Reader) (*BMPFile, error)
}

type BitmapParser struct {
	isParsed bool
}

func (p *BitmapParser) Parse(filepath string) (*BMPFile, error) {
	p.isParsed = true
	bmpFile := &BMPFile{}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// BMPHeader read
	err = binary.Read(file, binary.LittleEndian, &bmpFile.Header)
	if err != nil {
		return nil, err
	}

	// signature check
	if !isBM(bmpFile.Header.Type) {
		// TODO: исправить error:
		// return nil, fmt.Errorf("Error: %s is not bitmap file", filepath)
		return nil, errors.ErrNotBMPFile
	}

	// DIBHeader read
	err = binary.Read(file, binary.LittleEndian, &bmpFile.InfoHeader)
	if err != nil {
		return nil, err
	}

	// compression check
	if bmpFile.InfoHeader.Compression != 0 {
		return nil, errors.ErrCompressedBMP
	}

	// 24-bits check
	if bmpFile.InfoHeader.BitsPerPixel != 24 {
		return nil, errors.ErrUnsupportedBits
	}

	_, err = file.Seek(int64(bmpFile.Header.Offset), io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("error seeking to pixel data: %v", err)
	}

	// bitmap data read
	bmpFile.PixelData = make([]byte, bmpFile.InfoHeader.ImageSize)
	_, err = io.ReadFull(file, bmpFile.PixelData)
	if err != nil {
		return nil, fmt.Errorf("error reading pixel data: %v", err)
	}

	return bmpFile, nil
}

// isBM checks if the file type is BM and returns boolean.
func isBM(Type [2]byte) bool {
	return Type[0] == 'B' && Type[1] == 'M'
}
