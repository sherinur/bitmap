package main

import (
	"fmt"
	"os"

	"bitmap/internal/header"
	"bitmap/pkg"
)

func main() {
	args := os.Args[1:]
	if len(os.Args) < 2 {
		pkg.PrintUsage()
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "header":
		if len(args) <= 1 {
			pkg.PrintUsage()
			os.Exit(1)
		}

		err := header.Execute(args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", cmd)
		pkg.PrintUsage()
		os.Exit(1)
	}
}
