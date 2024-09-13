package bmp

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
