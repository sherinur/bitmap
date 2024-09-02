package pkg

import "fmt"

func PrintUsage(args ...string) {
	usage := `Usage: ./bitmap <command> [flags]`
	if len(args) == 1 {
		switch args[0] {
		case "header":
			usage = `Usage: ./bitmap header [filepath]`
		}
	}

	fmt.Println(usage)
}
