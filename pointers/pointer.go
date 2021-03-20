package main

import "fmt"

//pointer allows you to pass value types as references types
//"Am pointing to the value of the parameter"

var a int = 5
var c int = 6

func Zeroval(val int) {
	val = 0
	fmt.Println("val : ", val)
}

//set val to a param pointer
func Zeroptr(val *int) {
	*val = 0
	fmt.Println("pointer val : ", *val)
}

func printval() {
	s := 4
	b := &s
	fmt.Println(*b)
}

//using the new operator to create pointer
func newPointer() {
	a := new(int)
	*a = 4
	fmt.Println(*a)
}

//using ampersand to create a pointer
func ampersandPointer() {
	a := 2
	b := &a

	fmt.Println(*b)

}
