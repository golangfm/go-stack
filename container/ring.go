package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(2)
	fmt.Println(r.Len())
	var m map[int32]string
	fmt.Println(m[23])
}
