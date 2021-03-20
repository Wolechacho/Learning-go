package main

func callbackFunc(c, d int, f func(a, b int) int) int {
	return f(c, d)
}
