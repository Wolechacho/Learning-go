package main

func InitSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
