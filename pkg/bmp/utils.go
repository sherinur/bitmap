package bmp

import "fmt"

// isBM checks if the file type is BM and returns boolean.
func isBM(Type [2]byte) bool {
	return Type[0] == 'B' && Type[1] == 'M'
}

// convertToPixelArray converts the image data to two-dimensional pixel array.
func convertToPixelArray(data []byte, width int, height int) ([][]Pixel, error) {
	rowSize := (width*3 + 3) &^ 3
	expectedSize := rowSize * height

	if len(data) > expectedSize {
		data = data[:expectedSize]
	}

	if len(data) != expectedSize {
		return nil, ErrNotBMPFile
	}

	pixels := make([][]Pixel, height)
	for i := range pixels {
		pixels[i] = make([]Pixel, width)
	}

	for i := 0; i < height; i++ {
		rowStart := i * rowSize
		for j := 0; j < width; j++ {
			index := rowStart + j*3
			pixels[i][j] = Pixel{
				Blue:  data[index],
				Green: data[index+1],
				Red:   data[index+2],
			}
		}
	}

	return pixels, nil
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
