package main

func Sum(a int, b int) (int, error) {
	return a + b, nil
}

//variadic functions are like rest parameter in javascript

func Sum2(nums ...int) (int, error) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total, nil
}


