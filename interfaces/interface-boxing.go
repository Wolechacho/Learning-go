package main

import "fmt"

func Box() {
	var x interface{} = "foo"
	var s string = x.(string)
	s2, ok := x.(string)
	n, ok2 := x.(int)

	fmt.Println(s, s2, ok, n, ok2)
}
