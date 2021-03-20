package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func Shutdown() {
	panic(errors.New("Panic for panicking sake"))

	//defers are just like finally
}

//recover,defer,panic
func checkAndPrint(a []string, index int) {
	defer handleOutOfBounds()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}

//make sure the the recover is the defered function
func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
		fmt.Println("Stack Trace:")
		debug.PrintStack()
	}
}

//Return value of the function when panic is recovered
func checkAndGet(a []int, index int) (int, error) {
	defer handleOutOfBounds1()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	return a[index], nil
}

func handleOutOfBounds1() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}
