package main

import "fmt"

func main() {
	//less()
	//same()
	//variadicAppend()
	appendStringToByte()
}

//when the length < capacity
func less() {
	numbers := make([]int, 3, 5)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//Append number 4
	numbers = append(numbers, 4)
	fmt.Println("\nAppend Number 4")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//Append number 5
	numbers = append(numbers, 4)
	fmt.Println("\nAppend Number 5")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

//when the length == capacity .capacity multiply by 2
func same() {
	numbers := make([]int, 3, 3)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//Append number 4
	numbers = append(numbers, 4)
	fmt.Println("\nAppend Number 4")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

func variadicAppend() {
	numbers1 := []int{1, 2}
	numbers2 := []int{3, 4}
	numbers := append(numbers1, numbers2...)
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

//append function for string
func appendStringToByte() {
	sample := "Hello"
	suffix := " World"

	result := append([]byte(sample), suffix...)
	fmt.Printf("sample: %s\n", string(result))
}
