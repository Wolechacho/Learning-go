package main

import (
	"fmt"
)

var i = 1

func PrintLoop() {
	// for {
	// 	fmt.Println("Loop")
	// 	break
	// }

	for i := 2; i < 7; i++ {
		if i == 2 {
			fmt.Println(i)
		} else if i == 4 {
			fmt.Println(i)
		} else {
			continue
		}
	}

	j := 0

	switch j {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(2)
	default:
		fmt.Println("it is zero")
	}

	//A type switch compares types instead of values.
	//You can use this to discover the type of an interface value. In this example,
	//the variable t will have the type corresponding to its clause.

}

//LoopwithFxns - loop with a function has the initialization
func LoopwithFxns() {
	for test(); i < 3; i++ {
		fmt.Println(i)
	}
}

func test() {
	fmt.Println("testing functions")
}

//implement while loop
func while() {
	for i := 0; i < 5; i++ {
		if i < 5 {
			fmt.Println(i)
		}
	}
}
