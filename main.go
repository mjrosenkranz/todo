package main

import (
	"fmt"
	"github.com/xen0ne/todo/pkg/list"
)

func main() {
	// only top level lists
	lists := []*list.List{}

	lists = append(lists, list.NewList("the first list"))

	for _, l := range lists {
		fmt.Print(l.Print())
	}

}
