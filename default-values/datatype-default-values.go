package main

import (
	"fmt"
	"reflect"
)

//Pointy points
type Pointy struct {
	x int
	y int
}

var a int
var b string
var c bool
var d float64
var e byte
var x rune
var array [5]int
var sl []int
var keyValPair map[int]string
var channel chan string

type keet struct {
	F string `species:"gopher" color:"blue"`
	G string `species:"hopper" color:"green"`
}

//Learn - about default values of data types in golang
func Learn() {
	fmt.Println("Basic Data types")

	fmt.Printf("the value of a is %d\n", a)
	fmt.Printf("variable a data type is %s\n", reflect.TypeOf(a))

	fmt.Println()

	fmt.Printf("the value of b is %s\n", b)
	fmt.Printf("variable b data type is %s\n", reflect.TypeOf(b))

	fmt.Println()

	fmt.Printf("the value of c is %t\n", c)
	fmt.Printf("variable c data type is %s\n", reflect.TypeOf(c))

	fmt.Println()

	fmt.Printf("the value of d is %f\n", d)
	fmt.Printf("variable d data type is %s\n", reflect.TypeOf(d))

	fmt.Println()

	fmt.Printf("the value of e is %b\n", e)
	fmt.Printf("variable e data type is %s\n", reflect.TypeOf(e))

	fmt.Println()

	fmt.Printf("the value of x is %b\n", x)
	fmt.Printf("variable x data type is %s\n", reflect.TypeOf(x))

	fmt.Println()
	fmt.Println("End of Basic Data types")

	fmt.Println("Start of composite Data types")

	fmt.Println()
	fmt.Println("Start of reference data type under composite")

	fmt.Println()

	array[1] = 6
	fgh := array
	fgh[2] = 2
	fmt.Println("the value of array is ", array)
	fmt.Printf("variable array data type is %s\n", reflect.TypeOf(array))

	fmt.Println()

	fmt.Println("the value of fgh is ", fgh)
	fmt.Printf("variable fgh data type is %s\n", reflect.TypeOf(fgh))

	fmt.Println()

	p := Pointy{}
	fmt.Printf("the value of struct is %#v\n", p)
	fmt.Printf("variable struct data type is %s\n", reflect.TypeOf(p))

	ghk := sl
	//ghk[1] = 19
	sl = nil

	fmt.Println()

	fmt.Println("the value of sl is ", sl)
	fmt.Printf("variable sl data type is %s\n", reflect.TypeOf(sl))

	fmt.Println()

	fmt.Println("the value of ghk is ", ghk)
	fmt.Printf("variable ghk data type is %s\n", reflect.TypeOf(ghk))

	fmt.Println()

	fmt.Println("the value of map is ", keyValPair)
	fmt.Printf("variable keyValPair data type is %s\n", reflect.TypeOf(keyValPair))

	fmt.Println()

	fmt.Println("the value of channel is ", channel)
	fmt.Printf("variable channel data type is %s\n", reflect.TypeOf(channel))


}
