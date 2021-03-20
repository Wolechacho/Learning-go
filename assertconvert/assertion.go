package main

import (
	"fmt"
)

//AssertValue -assert
func AssertValue() {
	var i interface{}
	num := 5

	i = num

	val, _ := i.(int)

	fmt.Println(val)
}
