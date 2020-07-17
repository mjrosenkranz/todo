package main

import (
	"fmt"
	"github.com/xen0ne/todo/pkg/list"
)

func main() {
	l := list.NewList("list one")
	l2 := list.NewList("list bloop")

	l2.AddTask(list.NewTask("beep"))
	l2.AddTask(list.NewTask("bip"))

	l.AddTask(list.NewTask("task1"))
	l.AddTask(list.NewTask("task2"))
	l.AddSublist(l2)

	fmt.Print(l.Print())
}
