package main

import (
	"fmt"
)

var arr = make([]string, 0)

//Slices are reference types except you use the copy() function to clone the value only
func PrintSlices() {
	s := make([]string, 3)

	// Assign 1
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("Before")
	fmt.Println(s)

	//Assign 2
	fmt.Println("After")
	s = append(s, "d")
	fmt.Println(s)

	//Copy
	c := make([]string, len(s))
	//copy(c,s)
	c = s

	fmt.Println("Print C", c)

	c[0] = "z"
	s[1] = "y"
	fmt.Println("C after first element modification", c)
	fmt.Println("S after first element modification of c", s)

	//slice

	l := c[1:5]
	fmt.Println(l)

	g := c[2:]
	fmt.Println(g)

	h := c[:5]
	fmt.Println(h)

}

//copySlices - copy from a slice to another slice but in this instance the destination is preoccupied
func nonEmptyCopySlices() {
	s := []int{1, 2, 3, 8, 9}
	s1 := []int{4, 5, 6, 7}

	//specifies the number of element copied
	n2 := copy(s, s1)
	fmt.Println("Number of elements :", n2)

	fmt.Printf("s elem=%d,s=%v,s1=%v\n", len(s), s, s1)

	fmt.Println("s1 == nil", s1 == nil)

}

func copytonilemptySlices() {
	var s []int
	s1 := []int{4, 5, 6, 7}

	//specifies the number of element copied
	n2 := copy(s, s1)
	fmt.Println("Number of elements : ", n2)

	fmt.Printf("s elem=%d,s=%v,s1=%v\n", len(s), s, s1)
	fmt.Println("s == nil :", s == nil)

}

func sliceAll() {
	s1 := []int{4, 5, 6, 7}
	s1 = s1[:]
	fmt.Println(s1)
}

func clearSlicesWithoutAlloc() {
	s1 := []int{4, 5, 6, 7}
	s1 = s1[:0]
	fmt.Printf("len(s1) : %d, cap(s1) : %d , elem(s1) : %v", len(s1), cap(s1), s1)
}

func clearSlicesWitAlloc() {
	s1 := []int{4, 5, 6, 7}
	s1 = nil
	fmt.Printf("len(s1) : %d, cap(s1) : %d , elem(s1) : %v", len(s1), cap(s1), s1)
}

func exapmle2() (int, int) {
	return 1, 2
}

func nilSlices() {
	var sl3 []int
	sl3[0] = 3
	fmt.Println(sl3)
}

func acceptableSlice() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s []int
	s = a[2:4]

	fmt.Println(s)
	fmt.Println(a)
}

func exapmle3() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s []int
	s = a[2:4]

	s = append(s, 55, 66)
	s[0] = 33
	fmt.Println(s)
	fmt.Println(a)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
