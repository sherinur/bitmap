package main

import (
	"fmt"
	"os"

	"bitmap/internal/header"
	"bitmap/pkg"
)

func main() {
	if len(os.Args) < 2 {
		pkg.PrintUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "header":
		err := header.Execute(os.Args[2])
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
