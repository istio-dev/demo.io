package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	l.PushBack("cannon")

	l.PushFront(67)

	element := l.PushBack("first")

	l.InsertAfter("high", element)
	l.InsertBefore("noon", element)

	l.Remove(element)

	for i := l.Front(); i != nil ; i = i.Next() {
		fmt.Println(i.Value)
	}

}
