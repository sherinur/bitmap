package bmp

// parser.go provides functions to parse the bitmap(.bmp, .dib) file.

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// BitmapParser.Parse parses the bitmap(.bmp, .dib) file.
func (p *BitmapParser) Parse(filepath string) (*BMPFile, error) {
	// last moderated time check
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		p.isParsed = false
		return nil, err
	}

	if p.isParsed && p.lastParsedPath == filepath {
		if p.lastParsedBMP.ModTime.Before(fileInfo.ModTime()) {
			p.isParsed = false
		} else {
			return p.lastParsedBMP, nil
		}
	}

	bmpFile := &BMPFile{}
	// open file
	file, err := os.Open(filepath)
	if err != nil {
		p.isParsed = false
		return nil, err
	}
	defer file.Close()
	// BMPHeader read
	err = binary.Read(file, binary.LittleEndian, &bmpFile.Header)
	if err != nil {
		p.isParsed = false
		return nil, err
	}
	// signature check
	if !isBM(bmpFile.Header.Type) {
		return nil, fmt.Errorf("Error: %s is not bitmap file", filepath)
	}

	// DIBHeader read
	err = binary.Read(file, binary.LittleEndian, &bmpFile.InfoHeader)
	if err != nil {
		p.isParsed = false
		return nil, err
	}

	// compression check
	if bmpFile.InfoHeader.Compression != 0 {
		p.isParsed = false
		return nil, ErrCompressedBMP
	}

	// 24-bits check
	if bmpFile.InfoHeader.BitsPerPixel != 24 {
		p.isParsed = false
		return nil, ErrUnsupportedBits
	}

	// file info correction
	if bmpFile.InfoHeader.ImageSize == 0 {
		rowSize := (bmpFile.InfoHeader.Width*3 + 3) &^ 3
		bmpFile.InfoHeader.ImageSize = uint32(rowSize * bmpFile.InfoHeader.Height)
	}

	// bitmap data read
	_, err = file.Seek(int64(bmpFile.Header.Offset), io.SeekStart)
	if err != nil {
		p.isParsed = false
		return nil, err
	}

	data := make([]byte, bmpFile.InfoHeader.ImageSize)
	_, err = io.ReadFull(file, data)
	if err != nil {
		if err == io.EOF {
			p.isParsed = false
			return nil, fmt.Errorf("Error: File ended prematurely")
		}
		p.isParsed = false
		return nil, err
	}

	// // ! Debug print
	// fmt.Printf("Header Offset: %d, Width: %d, Height: %d, ImageSize: %d\n", bmpFile.Header.Offset, bmpFile.InfoHeader.Width, bmpFile.InfoHeader.Height, bmpFile.InfoHeader.ImageSize)

	// making two-dimensional slice of pixels
	bmpFile.ImageData, err = convertToPixelArray(data, int(bmpFile.InfoHeader.Width), int(bmpFile.InfoHeader.Height))
	if err != nil {
		p.isParsed = false
		return nil, err
	}

	p.isParsed = true
	p.lastParsedPath = filepath
	p.lastParsedBMP = bmpFile
	bmpFile.ModTime = fileInfo.ModTime()

	return bmpFile, nil
}
