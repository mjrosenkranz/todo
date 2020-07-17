package list

import "fmt"

const (
	max_level int = 3
	off_str string = "  "
)

type List struct {
	title string
	items []Item
	level int
}

func NewList(s string) *List {
	l := List{s, []Item{}, 0}
	return &l
}

func (l *List) set_level(lvl int) {
	l.level = lvl
}

func (l *List) AddSublist(a *List) {
	a.set_level(l.level + 1)
	l.items = append(l.items, a)
}

func (l *List) AddTask(t *Task) {
	l.items = append(l.items, t)
}

func (l List) Print() string {
	var ret string
	var offset string
	
	ret += fmt.Sprintf("%s\n", l.title)
	for i := 0; i < l.level + 1; i++ {
		offset += off_str
	}

	// ret += fmt.Sprintf("%s%s\n", offset, l.title)

	for _, itm := range l.items {
		ret += fmt.Sprintf("%s%s\n", offset, itm.Print())
	}

	return ret
}
