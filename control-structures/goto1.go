package main

import "fmt"

func learnGoto() {
	fmt.Println("a")
	goto FINISH
	fmt.Println("b")

FINISH:
	for i := 6; i < 10; i++ {
		fmt.Println(i,"e")
	}
	fmt.Println("c","e")
}

//compilation error
// func learnGoTo1() {
// 	fmt.Println("a")
// 	goto FINISH
// 	fmt.Println("b")

// }

// func test1() {
// FINISH:
// 	fmt.Println("c")
// }
