package bmp

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
