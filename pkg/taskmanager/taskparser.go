package taskmanager

import (
	"fmt"
	"os"
	"strings"

	"test/internal/tasks"
	"test/pkg/bmp"
)

func Parse(bmpFile *bmp.BMPFile, taskQueue *TaskQueue, options [][]string, startIndex int, endIndex int) error {
	for _, arg := range os.Args[startIndex:endIndex] {
		optionFound := false

		for _, option := range options {
			for _, flag := range option {
				// help flag check
				if strings.HasPrefix(arg, "--help") || strings.HasPrefix(arg, "-h") {
					return ErrHelpOccurs
				}

				formattedArg := "--" + flag + "="
				if strings.HasPrefix(arg, formattedArg) {
					optionFound = true
					value := strings.TrimPrefix(arg, formattedArg)
					newTask := NewTask()

					switch flag {
					case "mirror", "m":
						action := func(args ...string) {
							tasks.ApplyMirror(bmpFile, value)
							// fmt.Println("Mirror is applied! Options:", value)
						}
						newTask.SetAction(action)
					case "rotate", "r":
						action := func(args ...string) {
							tasks.ApplyRotate(bmpFile, value)
							// rotate.Execute(value)
						}
						newTask.SetAction(action)
					case "filter", "f":
						action := func(args ...string) {
							tasks.ApplyFilter(bmpFile, value)
							// fmt.Println("Filter is applied! Options:", value)
						}
						newTask.SetAction(action)
					case "crop", "c":
						action := func(args ...string) {
							tasks.ApplyCrop(bmpFile, value)
							// fmt.Println("Crop is applied! Options:", value)
						}
						newTask.SetAction(action)
					}

					taskQueue.Enqueue(newTask)
					break
				}
			}
		}

		if !optionFound {
			fmt.Printf("Undefined option: %s\n", arg)
			return ErrUndefinedOption
		}
	}

	return nil
}
