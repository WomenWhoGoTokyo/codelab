package taskmanager

import "fmt"

// Status is the progress of the task.
type Status int

const (
	// ToDo is the task progress status.
	ToDo Status = iota + 1
	// InProgress is the task progress status.
	InProgress
	// Done is the task progress status.
	Done
)

func (s Status) validate() error {
	switch s {
	case ToDo, InProgress, Done:
		return nil
	}
	return fmt.Errorf("%v is invalid Status.", s)
}
