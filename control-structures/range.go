package main

import (
	"fmt"
)

//range works on Array,slice,map,strings
func PrintRangeValues() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{6, 7, 8, 8, 10}
	word := "golang"
	keyValue := map[int]string{11: "eleven", 12: "tweleve", 13: "thirteen", 14: "fourteen", 15: "fifteen"}

	sum := 0
	sum1 := 0

	//array
	for _, num := range arr {
		sum += num
	}
	fmt.Println("Sum : ", sum)

	//slice
	for _, num := range slice {
		sum1 += num
	}
	fmt.Println("Sum : ", sum1)

	//strings
	for i, r := range word {
		fmt.Println(i, r)
	}

	//key value pair
	for key, value := range keyValue {
		fmt.Println(key, value)
	}
}

//With index and value
//With value only
//With index only
//Without index and value
func range2() {
	letters := []string{"a", "b", "c"}

	//With index and value
	fmt.Println("Both Index and Value")
	for i, letter := range letters {
		fmt.Printf("Index: %d Value:%s\n", i, letter)
	}

	//Only value
	fmt.Println("\nOnly value")
	for _, letter := range letters {
		fmt.Printf("Value: %s\n", letter)
	}

	//Only index
	fmt.Println("\nOnly Index")
	for i := range letters {
		fmt.Printf("Index: %d\n", i)
	}

	//Without index and value. Just print array values
	fmt.Println("\nWithout Index and Value")
	i := 0
	for range letters {
		fmt.Printf("Index: %d Value: %s\n", i, letters[i])
		i++
	}
}

//experimenting string using range and normal for loop
func loopThroughString() {
	sample := "a£c"
	fmt.Println(len(sample))

	//normal loop
	for i := 0; i < len(sample); i++ {
		fmt.Print(string(sample[i]))
	}

	fmt.Println()
	//range loop
	for _, letter := range sample {
		fmt.Print(string(letter))
	}

}
