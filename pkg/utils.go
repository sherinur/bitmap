package pkg

import "fmt"

func PrintUsage() {
	usage := `Usage: ./bitmap <command> [flags]`
	fmt.Println(usage)
}
