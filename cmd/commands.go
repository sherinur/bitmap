package cmd

import "fmt"

// commands.go stores data about commands: their Name, Description, Options and Usage.

type Command struct {
	Name        string
	Description string
	Usage       string
	Options     map[string]Option
}

type Option struct {
	Flag        []string
	Description string
}

// General option:
// "help": {
// 	Flag:        []string{"h", "help"},
// 	Description: "prints program usage information",
// },

var Commands = map[string]Command{
	"header": {
		Name:        "header",
		Description: "prints bitmap file header information",
		Usage:       "bitmap header <source_file>",
	},

	"apply": {
		Name:        "apply",
		Description: "applies processing to the image and saves it to the file",
		Usage:       "bitmap apply [options] <source_file> <output_file>",
		Options: map[string]Option{
			"rotate": {
				Flag:        []string{"r", "rotate"},
				Description: "rotates a bitmap image by a specified angle",
			},
			"filter": {
				Flag:        []string{"f", "filter"},
				Description: "applies various filters to image",
			},
			"crop": {
				Flag:        []string{"c", "crop"},
				Description: "trims a bitmap image according to specified parameters",
			},
			"mirror": {
				Flag:        []string{"m", "mirror"},
				Description: "mirrors a bitmap image either horizontally or vertically",
			},
		},
	},
}

func PrintUsage() {
	fmt.Println("Usage:\n  bitmap <command> [arguments]")
	fmt.Println("\nThe commands are:")

	alignment := findAlignment()

	for _, command := range Commands {
		fmt.Printf("  %-*s  %s\n", alignment, command.Name, command.Description)
	}
}

func PrintCommandHelp(commandName string) {
	if Commands[commandName].Usage != "" {
		usage := Commands[commandName].Usage
		fmt.Printf("Usage:\n  %s\n", usage)
	}
	if len(Commands[commandName].Options) > 0 {
		fmt.Println()
		PrintCommandOptions(commandName)
	}
}

func PrintCommandOptions(commandName string) {
	alignment := findAlignment()

	fmt.Println("The options are:")
	fmt.Printf("  %-*s  %s\n", alignment, "-h, --help", "prints program usage information")

	for _, option := range Commands[commandName].Options {
		var flags string
		for i, flag := range option.Flag {
			if i > 0 {
				flags += ", "
			}
			if len(flag) == 1 {
				flags += "-" + flag
			} else {
				flags += "--" + flag
			}
		}

		fmt.Printf("  %-*s  %s\n", alignment, flags, option.Description)
	}
}

// findAlignment finds the max string length that is used for alignment for print
func findAlignment() int {
	maxCommandLength := 0
	for _, command := range Commands {
		if len(command.Name) > maxCommandLength {
			maxCommandLength = len(command.Name)
		}
	}

	return maxCommandLength
}
