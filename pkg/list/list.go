package list

import "fmt"

type List struct {
	title string
	items []Item
}

func NewList(s string) List {
	l := List{s, make([]Item, 0)}
	return l
}

func (l *List) AddItem(i Item) {
	l.items = append(l.items, i)
}

func (l List) Print() string {
	var ret string
	ret += fmt.Sprintf("%s\n", l.title)
	for _, itm := range l.items {
		ret += fmt.Sprintf("%3s%s\n", " ", itm.Print())
	}

	return ret
}
