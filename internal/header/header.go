package header

import (
	"fmt"
<<<<<<< HEAD

	"bitmap/internal/bmp"
)

func Execute(filepath string) error {
	header, err := bmp.ExtractHeader(filepath)
	if err != nil {
		return err
	}

	fmt.Println(header)
=======
	"os"
)

func Execute(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// ? CONTINUE HERE ...
	// img, err := bmp.DecodeBMP(file)
	// if err != nil {
	// 	return err
	// }

	fmt.Println("Image decoded successfully.")
	// fmt.Println(img)
>>>>>>> dd64ac9fe426181de0efb39504f79a2887cf2365

	return nil
}
