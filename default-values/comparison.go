package main

import "fmt"

//compare - compare to similar data types togther, result should either be true or false
func compare() {
	//compare two struct

	type person struct {
		name string
		age  int
	}

	//consider the scenario where the struct fields have the different value
	p1 := person{name: "wole", age: 23}
	p2 := person{name: "bola", age: 21}
	fmt.Println("p1 and p2 : ", p1 == p2)

	//consider the scenario where the struct fields have the same value
	p3 := person{name: "tayo", age: 24}
	p4 := person{name: "tayo", age: 24}
	fmt.Println("p3 and p4 : ", p3 == p4)

	//Next is pointer using the struct example
	p5 := &person{name: "tayo", age: 24}
	p6 := &person{name: "tayo", age: 24}
	fmt.Println("pointers equal : ", p5 == p6)

	//Array
	//scenario they have different value
	var a [3]int = [3]int{1, 2, 3}
	var b [3]int = [3]int{1, 2, 4}

	fmt.Println("arrays with the different values : ", a == b)

	//scenario they have same value
	var c [3]int = [3]int{1, 2, 4}
	var d [3]int = [3]int{1, 2, 4}

	fmt.Println("arrays with the same values : ", c == d)

	//Slices
	//Illegal  : when you run it the compiler says invalid operation,slices can only be compared to nil

	// var s1 []int = []int{1,2,3}
	// var s2 []int = []int{1,2,4}

	// fmt.Println("slices equal : ", s1 == s2)

	//Map
	//Illegal  : when you run it the compiler says invalid operation,maps can only be compared to nil
	// var m1 map[int]string =  map[int]string{1:"one",2:"two"}
	// var m2 map[int]string =  map[int]string{1:"one",2:"two"}

	// fmt.Println("map equal : ", m1 == m2)

	//Function
	//Illegal  : when you run it the compiler says invalid operation,functions can only be compared to nil
	//fmt.Println(getNum == getNum1)

	fmt.Println()
	fmt.Println()
	//channels
	var ch chan int = make(chan int)
	var ch1 chan int = make(chan int)

	go func(a, b chan int) {
		fmt.Println(<-ch)
		fmt.Println(<-ch1)

		close(ch)
		close(ch1)
	}(ch, ch1)

	ch <- 2
	ch1 <- 2

	fmt.Println("ch1 and ch2 compared : ", ch == ch1)

	//Interfaces
	var f interface{} = 46
	var g interface{} = 47
	//interface with the same data type but different value
	fmt.Println("interface with the same data type  but different value : ", f == g)

	var f1 interface{} = 46
	var g1 interface{} = 46
	//interface with the same data type and the same value
	fmt.Println("interface with the same data type and the same value : ", f1 == g1)

	var f2 interface{} = 46
	var g2 interface{} = "word"
	//interface with the same different type
	fmt.Println("interface with the different type  : ", f2 == g2)

}

func getNum() {
	return
}

func getNum1() {
	return
}
