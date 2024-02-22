package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Mohamad")
	data.PushBack("Ali")
	data.PushBack("Farhan")

	var head *list.Element = data.Front()
	head.Next()

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
