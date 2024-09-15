package apply

import (
	"fmt"
	"os"

	"bitmap/cmd"
	"bitmap/pkg/bmp"
	"bitmap/pkg/taskmanager"
)

func Execute(newFilepath string, taskQueue *taskmanager.TaskQueue, bmpFile *bmp.BMPFile) error {
	// creating options and their actions
	var options [][]string
	for _, option := range cmd.Commands["apply"].Options {
		options = append(options, option.Flag)
	}

	// parsing flags and values, filling task queue
	err := taskmanager.Parse(bmpFile, taskQueue, options, 2, len(os.Args)-2)
	if err != nil {
		if err == taskmanager.ErrHelpOccurs || err == taskmanager.ErrUndefinedOption {
			cmd.PrintCommandHelp("apply")
			os.Exit(1)
		}

		fmt.Println(err)
		os.Exit(1)
	}

	if bmpFile == nil {
		fmt.Println("bmpFile is nil after parsing!")
		return fmt.Errorf("Failed to parse BMP file")
	}

	// handling taskqueue
	err = taskmanager.Handler(taskQueue)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bmp.SaveBMP(newFilepath, bmpFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// bmpFile.DebugPrint()

	return nil
}
