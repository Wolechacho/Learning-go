package main

import "fmt"

//breaking or not breaking rule
// const (
// 	C0 = iota
// 	C1 = iota
// 	C2 = iota
// ) 0,1,2

// const (
// 	C0 = iota
// 	C1
// 	C2
// )0,1,2

// const (
// 	C0 = 1 + iota
// 	C1
// 	C2
// )1,2,3

// const (
// 	C0 = 1 + iota
// 	C1 = iota
// 	C2 = iota
// )1,1,2

const (
	C0 = 1 + iota
	_
	C1 = iota
	C2 = iota
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func PrintIotaVals() {
	fmt.Println(C0, C1, C2)
}
