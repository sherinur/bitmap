package main

import (
	"fmt"
	"os"

	"bitmap/internal/header"
	"bitmap/internal/mirror"
	"bitmap/pkg/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		utils.PrintUsage()
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "header":
		if len(args) <= 1 {
			utils.PrintUsage("header")
			os.Exit(1)
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
	case "mirror":
		if len(args) < 3 {
			utils.PrintUsage()
			os.Exit(1)
		}

		err := mirror.Execute(args[1], args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", cmd)
		utils.PrintUsage()
		os.Exit(1)
	}
}
