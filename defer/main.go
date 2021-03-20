package main

import "fmt"

func main() {
	//defer test()
	//fmt.Println("Executed in main")

	//evaluation of defer function with arguments
	//deferedArgs()

	//multiple defer functions in the same function
	//dfunc()

	//Defer function and Named Return Values
	 //s := deferedNamedFunc()
	//fmt.Println(s)

	//defer with panic
	//deferPanic()

	//anonymous function
	deferAnonFunc()

	//embedded defered functions
	f1()
}
func test() {
	fmt.Println("In Defer")
}

//anonymous function defered
func deferAnonFunc() {
	defer func() {
		fmt.Println("Anonymous function defered")
	}()
	fmt.Println("hello")
}

//evaluation of defer function with arguments
func deferedArgs() {
	sample := "abc"

	defer fmt.Printf("In defer sample is: %s\n", sample)
	sample = "xyz"
}

//multiple defer functions in the same function - LIFO
func dfunc() {
	i := 0
	i = 1
	defer fmt.Println(i)
	i = 2
	defer fmt.Println(i)
	i = 3
	defer fmt.Println(i)
}

//Defer function and Named Return Values
func deferedNamedFunc() (size int) {
	defer func() { size = 20 }()
	size = 30
	return
}

//defer with panic
func deferPanic() {
	defer fmt.Println("Defer in main")
	panic("Panic with Defer")
	fmt.Println("After painc in f2")
}

func f1() {
	fmt.Println("Hello from func 1")
	defer func() { fmt.Println("this function is defered from f1()") }()
	f2()
}

func f2() {
	fmt.Println("Hello from func 2")
	defer func() { fmt.Println("this function is defered from f2()") }()
	f3()

}
func f3() {
	fmt.Println("Hello from func 3")
	defer func() { fmt.Println("this function is defered from f3()") }()
}
