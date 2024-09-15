package main

import (
	"fmt"
	"os"

	"bitmap/cmd"
	"bitmap/internal/apply"
	"bitmap/internal/header"
	"bitmap/pkg/bmp"
	"bitmap/pkg/taskmanager"
)

var (
	// global variables
	GlobalBmpFile   *bmp.BMPFile
	GlobalTaskQueue = taskmanager.NewTaskQueue()
)

func main() {
	if len(os.Args) <= 1 {
		cmd.PrintUsage()
		os.Exit(1)
	}
	args := os.Args[1:]
	command := args[0]

	switch command {
	case "apply":
		if len(os.Args) <= 2 {
			cmd.PrintCommandHelp(command)
			os.Exit(1)
		}

		lastArg := os.Args[len(os.Args)-1]
		secondToLastArg := args[len(args)-2]
		var newFilepath, filepath string
		if lastArg != "-h" && lastArg != "--help" && secondToLastArg != "-h" && secondToLastArg != "--help" {
			newFilepath = lastArg
			filepath = secondToLastArg
		} else {
			cmd.PrintCommandHelp(command)
			os.Exit(1)
		}

		// parse and store bmpfile
		p := bmp.BitmapParser{}
		GlobalBmpFile, err := p.Parse(filepath)
		if err != nil {
			fmt.Println(err)
			cmd.PrintCommandHelp(command)
			os.Exit(1)
		}

		err = apply.Execute(newFilepath, &GlobalTaskQueue, GlobalBmpFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "header":
		if len(os.Args) <= 2 {
			cmd.PrintCommandHelp(command)
			os.Exit(1)
		}

		for _, arg := range args[1:] {
			if arg == "-h" || arg == "--help" {
				cmd.PrintCommandHelp(command)
				os.Exit(1)
			}
		}

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
