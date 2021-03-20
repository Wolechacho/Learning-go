package main

import "fmt"

var t = 256

func main() {
	compare()
	//variables()
	//Learn()
}

func variables() {
	var t = 90
	fmt.Println(t)

	//example 1 - single declaration
	var a = 3
	fmt.Println(a)

	//example 2 - multiple declaration
	var b, c = 4, 90
	fmt.Println(b)
	fmt.Println(c)

	//short single declaration
	d := true
	fmt.Println(d)

	//short multiple declaration
	e, f := true, 1
	fmt.Println(e)
	fmt.Println(f)

	//hacky - this only possible if one of the variable at the left side is new
	//if you set it to another data type there is going to be a compiler error
	g, h := 1, 2
	h, j := 5, 4
	fmt.Println("Hacky", g, h, j)

}
