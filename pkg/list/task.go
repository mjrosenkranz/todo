package list

import (
	"fmt"
)

const (
	comp_str= "[ ]"
	incomp_str= "[x]"
)

type Task struct {
	task string
	completed bool
}

func NewTask(s string) *Task {
	return &Task{s, false}
}

func (t *Task) Toggle() {
	fmt.Println("yeet")
	t.completed = !t.completed
}


func (t Task) Print() string {
	if t.completed  {
		return fmt.Sprintf("%s %s", comp_str, t.task)
	} else {
		return fmt.Sprintf("%s %s", incomp_str, t.task)
	}
}
