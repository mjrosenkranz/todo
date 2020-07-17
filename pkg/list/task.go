package list

import (
	"fmt"
)

type Task struct {
	task string
}

func NewTask(s string) *Task {
	return &Task{s}
}

func (t Task) Print() string {
	return fmt.Sprintf("%s %s", "[ ]", "t.task")
}
