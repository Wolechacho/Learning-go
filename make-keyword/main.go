package main

import "fmt"

func main() {
	var a []int

	var b []int = make([]int, 2)

	//a = nil
	b = nil
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))

	//illegal b is already set to be zero
	b[0] = 4
	if a == nil{
		fmt.Println("a is nil")
	}
}
