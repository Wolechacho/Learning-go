package main

import (
	"fmt"
	"mygoplayground/learn-packages/math"
)

func init() {
	fmt.Println("In main init")
}

func main() {
	fmt.Println(math.Add(2, 1))
	fmt.Println(math.Subtract(2, 1))
}
