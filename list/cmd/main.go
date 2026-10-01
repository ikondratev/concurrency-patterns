package main

import (
	"fmt"

	"github.com/kilia/list-example/list"
)

func printResult(l list.Forward[int]) {
	for i := l; i != nil; i = i.Next() {
		fmt.Println(i.Data())
	}
}

func main() {
	l := list.NewTwList(5)
	d := list.NewOwList(3)

	l.Push(7)
	l.Push(9)
	l = l.PushPop(3)
	printResult(l)
	printResult(d)
}