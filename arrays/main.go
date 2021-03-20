package main

import (
	"fmt"
)

//Note Arrays are not reference type unlike in some programming language
func PrintArray() {
	//you can set an array like this ,use make for slices,map and channels
	var numbers [5]int

	numbers[2] = 20
	numbers[4] = 1

	// for i:= 0 ; i < len(numbers);i++{
	// 		fmt.Println(numbers[i])
	// }

	num2 := numbers

	//you can slice out an arrray
	num3 := numbers[1:3]
	fmt.Println("num3 : ", num3)

	fmt.Println("Before")
	fmt.Println(numbers)
	fmt.Println(num2)

	numbers[0] = 1
	num2[2] = 80

	fmt.Println("After")
	fmt.Println(numbers)
	fmt.Println(num2)
}

func main() {
	PrintArray()
}
