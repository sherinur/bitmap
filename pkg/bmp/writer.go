package bmp

import (
	"encoding/binary"
	"os"
)

// bmp.SaveBMP saves the image data to a BMP file
func SaveBMP(filename string, bmpFile *BMPFile) error {
	// if !validateBMPData(bmpFile) {
	// 	return ErrSaveError
	// }

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = writeBMPHeader(file, &bmpFile.Header)
	if err != nil {
		return err
	}

	err = writeDIBHeader(file, &bmpFile.InfoHeader)
	if err != nil {
		return err
	}

	err = writeImageData(file, bmpFile.ImageData)
	if err != nil {
		return err
	}

	filepadding := int(bmpFile.InfoHeader.ImageSize) % 4
	for i := 0; i < filepadding; i++ {
		err := binary.Write(file, binary.LittleEndian, byte(0))
		if err != nil {
			return err
		}
	}

	return nil
}

// writeImageData writes the image data to the file
func writeImageData(file *os.File, imageData [][]Pixel) error {
	bytesPerPixel := 3

	for _, row := range imageData {
		// Write pixels
		for _, pixel := range row {
			err := binary.Write(file, binary.LittleEndian, pixel)
			if err != nil {
				return err
			}
		}

		// Calculate padding for the row
		rowSize := len(row) * bytesPerPixel
		padding := (4 - (rowSize % 4)) % 4

		// Write padding bytes
		for i := 0; i < padding; i++ {
			err := binary.Write(file, binary.LittleEndian, byte(0))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// writeBMPHeader writes the BMP file header to the file
func writeBMPHeader(file *os.File, header *BMPHeader) error {
	dataToWrite := make([]byte, 0, 14)
	dataToWrite = append(dataToWrite, header.Type[:]...)

	buf := make([]byte, 4)

	binary.LittleEndian.PutUint32(buf, header.Size)
	dataToWrite = append(dataToWrite, buf...)

	buf = make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, header.Reserved1)
	dataToWrite = append(dataToWrite, buf...)

	binary.LittleEndian.PutUint16(buf, header.Reserved2)
	dataToWrite = append(dataToWrite, buf...)

	buf = make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, header.Offset)
	dataToWrite = append(dataToWrite, buf...)

	_, err := file.Write(dataToWrite)
	if err != nil {
		return err
	}

	return nil
}

// writeDIBHeader writes the DIB header to the file
func writeDIBHeader(file *os.File, infoHeader *DIBHeader) error {
	dataToWrite := make([]byte, 0, 40)

	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, infoHeader.Size)
	dataToWrite = append(dataToWrite, buf...)

	binary.LittleEndian.PutUint32(buf, uint32(infoHeader.Width))
	dataToWrite = append(dataToWrite, buf...)

	binary.LittleEndian.PutUint32(buf, uint32(infoHeader.Height))
	dataToWrite = append(dataToWrite, buf...)

	binary.LittleEndian.PutUint16(buf[:2], infoHeader.Planes)
	dataToWrite = append(dataToWrite, buf[:2]...)

	binary.LittleEndian.PutUint16(buf[:2], infoHeader.BitsPerPixel)
	dataToWrite = append(dataToWrite, buf[:2]...)

	binary.LittleEndian.PutUint32(buf, infoHeader.Compression)
	dataToWrite = append(dataToWrite, buf...)

	binary.LittleEndian.PutUint32(buf, infoHeader.ImageSize)
	dataToWrite = append(dataToWrite, buf...)

	binary.LittleEndian.PutUint32(buf, uint32(infoHeader.XPixelsPerMeter))
	dataToWrite = append(dataToWrite, buf...)

	binary.LittleEndian.PutUint32(buf, uint32(infoHeader.YPixelsPerMeter))
	dataToWrite = append(dataToWrite, buf...)

	binary.LittleEndian.PutUint32(buf, infoHeader.ColorsUsed)
	dataToWrite = append(dataToWrite, buf...)

	binary.LittleEndian.PutUint32(buf, infoHeader.ColorsImportant)
	dataToWrite = append(dataToWrite, buf...)

	_, err := file.Write(dataToWrite)
	if err != nil {
		return err
	}

	return nil
}

// // bmp.ValidateBMPData() checks if the BMP data is valid
// func validateBMPData(bmpFile *BMPFile) bool {
// 	// TODO: implement validating bmp
// 	return true
// }
