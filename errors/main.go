package main

import "fmt"

func main() {
	//sampError()
	//formatError()
	// err := validate("wole", "")
	// if err != nil {
	// 	if err, ok := err.(*inputError); ok {
	// 		fmt.Println(err)
	// 		fmt.Printf("Missing Field is %s\n", err.getMissingField())
	// 	}
	// }

	//wrap error
	//wrapError()

	//unwrap error
	//unwrapError()

	//wrap error2
	// num := -1
	// err := checkPostiveAndEven(num)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Given number is positive and even")
	// }

	//check if two errors are equal
	// var err1 errorOne
	// err2 := do()
	// if err1 == err2 {
	// 	fmt.Println("Equality Operator: Both errors are equal")
	// }
	// if errors.Is(err1, err2) {
	// 	fmt.Println("Is function: Both errors are equal")
	// }

	//why you should use errors.Is for comparison
	// err1 := errorOne{}

	// err2 := do1()

	// if err1 == err2 {
	// 	fmt.Println("Equality Operator: Both errors are equal")
	// } else {
	// 	fmt.Println("Equality Operator: Both errors are not equal")
	// }

	// if errors.Is(err2, err1) {
	// 	fmt.Println("Is function: Both errors are equal")
	// }

	//recover
	// a := []string{"a", "b"}
	// checkAndPrint(a, 2)
	// fmt.Println("Exiting normally")

	//Return value of the function when panic is recovered
	a := []int{5, 6}
	val, err := checkAndGet(a, 2)
	fmt.Printf("Val: %d\n", val)
	fmt.Println("Error: ", err)
}
