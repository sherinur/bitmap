package taskmanager

import (
	"fmt"
	"os"
	"strings"

	"bitmap/internal/tasks"
	"bitmap/pkg/bmp"
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
				longFlag := "--" + flag + "="
				shortFlag := "-" + flag + "="

				if strings.HasPrefix(arg, longFlag) || strings.HasPrefix(arg, shortFlag) {
					optionFound = true
					value := strings.TrimPrefix(arg, longFlag)
					if value == arg {
						value = strings.TrimPrefix(arg, shortFlag)
					}
					newTask := NewTask()

					switch flag {
					case "mirror", "m":
						action := func(args ...string) {
							tasks.ApplyMirror(bmpFile, value)
						}
						newTask.SetAction(action)
					case "rotate", "r":
						action := func(args ...string) {
							tasks.ApplyRotate(bmpFile, value)
						}
						newTask.SetAction(action)
					case "filter", "f":
						action := func(args ...string) {
							tasks.ApplyFilter(bmpFile, value)
						}
						newTask.SetAction(action)
					case "crop", "c":
						action := func(args ...string) {
							tasks.ApplyCrop(bmpFile, value)
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
