package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l list.List
	m := l.PushBack("abc")
	m1 := l.InsertAfter(123, m)
	l.InsertBefore("ttt", m1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println(l.Len())
}
