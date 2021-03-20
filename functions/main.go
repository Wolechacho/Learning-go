package main

import "fmt"

func main() {
	a := 2
	b := 3

	g := callbackFunc(a, b, func(h,i int)int{
		return h + i
	})
	fmt.Println(g)

	init := InitSeq()
	init()
	init()
	u := init()
	fmt.Println(u)
	
}

func sum(g, h int) int {
	return g + h
}
