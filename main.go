package main

import (
	"fmt"
	"os"

	"test/cmd"
	"test/internal/apply"
	"test/internal/header"
	"test/pkg/bmp"
	"test/pkg/taskmanager"
)

var (
	// global variables
	GlobalBmpFile   *bmp.BMPFile
	GlobalTaskQueue = taskmanager.NewTaskQueue()
)

func main() {
	if len(os.Args) <= 2 {
		cmd.PrintUsage()
		os.Exit(1)
	}
	args := os.Args[1:]
	command := args[0]

	switch command {
	case "apply":
		newFilepath := args[len(args)-1]
		filepath := args[len(args)-2]

		// parse and store bmpfile
		p := bmp.BitmapParser{}
		GlobalBmpFile, err := p.Parse(filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = apply.Execute(newFilepath, &GlobalTaskQueue, GlobalBmpFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "header":
		for i := 1; i < len(args); i++ {
			err := header.Execute(args[i])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if i != len(args)-1 {
				// newline after each info, except the last one
				fmt.Println()
			}
		}
	default:
		cmd.PrintUsage()
		os.Exit(1)
	}
}
