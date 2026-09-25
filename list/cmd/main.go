package main

import (
	"fmt"

	"github.com/kilia/list-example/list"
)

func main() {
	l := list.NewTwList(5)

	l.Push(7)
	l.Push(9)
	l = l.PushPop(3)

	for i := l; i != nil; i = i.Next {
		fmt.Println(i.Data)
	}
}