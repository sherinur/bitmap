package taskmanager

import "errors"

var (
	ErrQueueEmpty = errors.New("Tasks error: Queue is empty")
	ErrHelpOccurs = errors.New("Tasks error: Help occurs")

	ErrUndefinedOption = errors.New("Error: Undefined option")
	ErrUndefinedValue  = errors.New("Error: Undefined value")
)
