package bmp

import "fmt"

// applying negative filter
func (b *BMPFile) ApplyNegative() {
	grid := b.ImageData

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j].Red = 255 - grid[i][j].Red
			grid[i][j].Green = 255 - grid[i][j].Green
			grid[i][j].Blue = 255 - grid[i][j].Blue
		}
	}
}

// applying grayscale
func (b *BMPFile) ApplyGrayScale() {
	grid := b.ImageData

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			avg := grid[i][j].Red + grid[i][j].Green + grid[i][j].Blue
			grid[i][j].Red = avg
			grid[i][j].Blue = avg
			grid[i][j].Green = avg
		}
	}
}

// applying filter by color
func (bmpfile *BMPFile) ApplyByColor(color string) {
	grid := bmpfile.ImageData

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch color {
			case "red":
				grid[i][j].Blue = 0
				grid[i][j].Green = 0
			case "blue":
				grid[i][j].Green = 0
				grid[i][j].Red = 0
			case "green":
				grid[i][j].Blue = 0
				grid[i][j].Red = 0
			}
		}
	}
}

// applying blur filter
func (b *BMPFile) ApplyBlur() {
	grid := b.ImageData
	width := len(grid[0])
	height := len(grid)

	blurredGrid := make([][]Pixel, height)
	for i := range blurredGrid {
		blurredGrid[i] = make([]Pixel, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var sumRed, sumGreen, sumBlue int
			count := 1

			for dx := -10; dx <= 10; dx++ {
				for dy := -10; dy <= 10; dy++ {
					nx, ny := i+dx, j+dy

					if nx >= 0 && nx < height && ny >= 0 && ny < width {
						sumRed += int(grid[nx][ny].Red)
						sumGreen += int(grid[nx][ny].Green)
						sumBlue += int(grid[nx][ny].Blue)
						count++
					}
				}
			}

			blurredGrid[i][j].Red = uint8(sumRed / count)
			blurredGrid[i][j].Green = uint8(sumGreen / count)
			blurredGrid[i][j].Blue = uint8(sumBlue / count)
		}
	}

	b.ImageData = blurredGrid
}

// apply pixelate function with the blocksize(default:20)
func (b *BMPFile) ApplyPixelation(blockSize int) {
	grid := b.ImageData
	width := len(grid[0])
	height := len(grid)

	for i := 0; i < height; i += blockSize {
		for j := 0; j < width; j += blockSize {
			var sumRed, sumGreen, sumBlue int
			var count int
			for dx := 0; dx < blockSize; dx++ {
				for dy := 0; dy < blockSize; dy++ {
					nx, ny := dx+i, dy+j
					if nx < height && ny < width {
						sumGreen += int(grid[nx][ny].Green)
						sumRed += int(grid[nx][ny].Red)
						sumBlue += int(grid[nx][ny].Blue)
						count++
					}
				}
			}
			avgRed := uint8(sumRed / count)
			avgGreen := uint8(sumGreen / count)
			avgBlue := uint8(sumBlue / count)
			for dx := 0; dx < blockSize; dx++ {
				for dy := 0; dy < blockSize; dy++ {
					nx, ny := dx+i, dy+j
					if nx < height && ny < width {
						grid[nx][ny].Green = avgGreen
						grid[nx][ny].Red = avgRed
						grid[nx][ny].Blue = avgBlue
					}
				}
			}
		}
	}
}

// cropImage() performs the actual cropping of the BMP image
func (bmpFile *BMPFile) CropImage(OffsetX, OffsetY, Width, Height int) {
	croppedData := make([][]Pixel, Height)
	for i := 0; i < Height; i++ {
		croppedData[i] = make([]Pixel, Width)
		copy(croppedData[i], bmpFile.ImageData[int(bmpFile.InfoHeader.Height)-Height-OffsetY+i][OffsetX:OffsetX+Width])
	}

	// calculating Iamge size and File size in bytes
	NewIamgesize := ((Width*3 + 3) &^ 3) * Height
	NewFileSizeinBytes := NewIamgesize + 54
	// update image data and DIB header and Header
	bmpFile.ImageData = croppedData
	bmpFile.InfoHeader.Width = int32(Width)
	bmpFile.InfoHeader.Height = int32(Height)
	bmpFile.InfoHeader.ImageSize = uint32(NewIamgesize)
	bmpFile.Header.Size = uint32(NewFileSizeinBytes)
}

// ApplyMirrorHorizontal mirrors the image vertically by swapping pixels within each column.
func (bmpFile *BMPFile) ApplyMirrorVertical() error {
	rowsNum := len(bmpFile.ImageData)
	if rowsNum == 0 {
		return nil // No rows to process
	}

	for i := 0; i < rowsNum/2; i++ {
		// Calculate the index of the row to swap with
		oppRow := rowsNum - 1 - i
		// Swap the rows
		bmpFile.ImageData[i], bmpFile.ImageData[oppRow] = bmpFile.ImageData[oppRow], bmpFile.ImageData[i]
	}

	return nil
}

// ApplyMirrorHorizontal mirrors the image horizontally by swapping pixels within each row.
func (bmpFile *BMPFile) ApplyMirrorHorizontal() error {
	rowsNum := len(bmpFile.ImageData)
	if rowsNum == 0 {
		return nil // No rows to process
	}

	colsNum := len(bmpFile.ImageData[0])
	for i := 0; i < rowsNum; i++ {
		// Get the current row
		row := bmpFile.ImageData[i]
		// Swap pixels horizontally within the row
		for j := 0; j < colsNum/2; j++ {
			oppCol := colsNum - 1 - j
			// Swap pixels
			row[j], row[oppCol] = row[oppCol], row[j]
		}
	}

	return nil
}

// ApplyRotation rotates the BMP file by specified angles: 90, -90, 180, 270, or -270 degrees.
func (bmpFile *BMPFile) ApplyRotation(angle string) error {
	switch angle {
	case "-90", "270", "left":
		bmpFile.ImageData = rotate90(bmpFile)
		// Swap width and height
		bmpFile.InfoHeader.Width, bmpFile.InfoHeader.Height = bmpFile.InfoHeader.Height, bmpFile.InfoHeader.Width
	case "90", "-270", "right":
		bmpFile.ImageData = rotate90CounterClockwise(bmpFile)

		bmpFile.InfoHeader.Width, bmpFile.InfoHeader.Height = bmpFile.InfoHeader.Height, bmpFile.InfoHeader.Width
	case "180", "-180":
		bmpFile.ImageData = rotate180(bmpFile)
	default:
		return fmt.Errorf("unsupported rotation angle: %s", angle)
	}

	// Update ImageSize after rotation
	bytesPerPixel := int(bmpFile.InfoHeader.BitsPerPixel / 8)
	rowSize := ((int(bmpFile.InfoHeader.Width) * bytesPerPixel) + 3) &^ 3 // including padding
	bmpFile.InfoHeader.ImageSize = uint32(rowSize * int(bmpFile.InfoHeader.Height))

	// Update BMP header size
	bmpFile.Header.Size = 54 + bmpFile.InfoHeader.ImageSize

	return nil
}

// rotate90 rotates the image data by 90 degrees clockwise.
func rotate90(bmpFile *BMPFile) [][]Pixel {
	width := len(bmpFile.ImageData[0])
	height := len(bmpFile.ImageData)
	newData := make([][]Pixel, width)

	for i := range newData {
		newData[i] = make([]Pixel, height)
		for j := range newData[i] {
			newData[i][j] = bmpFile.ImageData[height-1-j][i]
		}
	}

	return newData
}

// rotate90CounterClockwise rotates the image data by 90 degrees counter-clockwise.
func rotate90CounterClockwise(bmpFile *BMPFile) [][]Pixel {
	width := len(bmpFile.ImageData[0])
	height := len(bmpFile.ImageData)
	newData := make([][]Pixel, width)

	for i := range newData {
		newData[i] = make([]Pixel, height)
		for j := range newData[i] {
			newData[i][j] = bmpFile.ImageData[j][width-1-i]
		}
	}

	return newData
}

// rotate180 rotates the image data by 180 degrees.
func rotate180(bmpFile *BMPFile) [][]Pixel {
	width := len(bmpFile.ImageData[0])
	height := len(bmpFile.ImageData)
	newData := make([][]Pixel, height)

	for i := range newData {
		newData[i] = make([]Pixel, width)
		for j := range newData[i] {
			newData[i][j] = bmpFile.ImageData[height-1-i][width-1-j]
		}
	}

	return newData
}

// for debugging
func (bmpFile *BMPFile) DebugPrint() {
	// fmt.Printf("Header Offset: %d, Width: %d, Height: %d, ImageSize: %d\n", bmpFile.Header.Offset, bmpFile.InfoHeader.Width, bmpFile.InfoHeader.Height, bmpFile.InfoHeader.ImageSize)
	fmt.Println(len(bmpFile.ImageData))
}
