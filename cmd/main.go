package main

import (
	"fmt"
	"os"

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
		fmt.Println(cmd)
	default:
		fmt.Println("Unknown command.")
		os.Exit(1)
	}
}
