package main

import (
	"fmt"
)

//Maps reference types
func PrintMap() {
	s := make(map[int]string)

	s[1] = "one"
	s[2] = "two"
	s[3] = "three"
	fmt.Println(s)

	delete(s, 1)

	fmt.Println(s)

	//check if key is present
	_, isKey := s[2]
	fmt.Println(isKey)

	//declare and initialize
	//n := map[string]string{"1": "one", "2": "two"}

}
