package main

import (
	"fmt"
	"github.com/xen0ne/todo/pkg/list"
)

func main() {
	l := list.NewList("list one")
	l2 := list.NewList("list 2")

	l2.AddItem(list.NewTask("task2"))
	l2.AddItem(list.NewTask("task1"))
	l2.AddItem(list.NewTask("task2"))

	l.AddItem(list.NewTask("task1"))
	l.AddItem(list.NewTask("task2"))
	l.AddItem(l2)

	fmt.Print(l.Print())
}
