package main

import (
	"fmt"
	"math"
)

const (
	t = "boy"
	g = "girl"
)
const aa = 123

var (
	toddle = "toddle"
	dober  = "dobberman"
)

func PrintConstants() {
	fmt.Println(t, g)

	//multiple declaration
	const a, b = 1, 5
	fmt.Println(a, b)

	//error of redeclaring const
	// const test = "name"
	// test = "test"

	//error of assigning runtime executable function
	//const initializer getValue() is not a constant
	//const val = getValue()

	//test for constant declared as global and local with the same variable name
	const aa = 456
	fmt.Println(aa)

}

func getValue() int {
	return 1
}

//setting type and untyped constant to a variable
func TypedConstants() {
	//typed constants
	const a int32 = 8

	//var i int32
	//var j int64

	//this works fine
	// i = a
	// fmt.Println(i)
	//this will flag error
	//j = a

	//untyped constants
	var f1 float32
	var f2 float64
	//math.PI is an untyped constant so it can be assigned to other types
	f1 = math.Pi
	f2 = math.Pi

	fmt.Printf("Type: %T Value: %v\n", math.Pi, math.Pi)
	fmt.Printf("Type: %T Value: %v\n", f1, f1)
	fmt.Printf("Type: %T Value: %v\n", f2, f2)

	//typed constant
	type myString string
	const s string = "sole"
	b := s
	fmt.Println(b)

	//error
	//var f myString = s

	//untyped constant
	type myString1 string
	const v = "sole"
	h := v
	fmt.Println(h)

	var n myString1 = v
	fmt.Println(n)

}

func UnsupportedConstantTypes() {
	//map error
	// const mp = map[string]int{
	// 	"one": 1,
	// }

	//struct error
	// const e = employee{
	// 	name: "John",
	// 	age:  21,
	// }
	// fmt.Println(e)

	//slice error
	//const e = [1]int{1}

	//interface error
	//const s interface{} = "wole"

}

type employee struct {
	name string
	age  int
}
