package main

import (
	"fmt"
	"runtime"
)

func main() {
	//printval()
	// Zeroval(a)
	// fmt.Println("a : ",a)

	// Zeroptr(&c)
	// fmt.Println("c : ",c)

	//Declare
	// var b *int
	// a := 2
	// b = &a

	//Will print a address. Output will be different everytime.
	// fmt.Println(b)
	// fmt.Println(*b)
	// b = new(int)
	// *b = 10
	// fmt.Println(*b)

	fmt.Println(runtime.NumCPU())
}
