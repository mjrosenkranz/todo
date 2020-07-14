package list

type Task struct {
	task string
}

func NewTask(s string) Task {
	return Task{s}
}

func (t Task) Print() string {
	return t.task
}
