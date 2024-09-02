package main

import (
	"fmt"
	"os"

	"bitmap/internal/header"
	"bitmap/pkg"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		pkg.PrintUsage()
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "header":
		if len(args) <= 1 {
			pkg.PrintUsage("header")
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
	default:
		fmt.Println("Unknown command:", cmd)
		pkg.PrintUsage()
		os.Exit(1)
	}
}
