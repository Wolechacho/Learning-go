package main

import "fmt"

type dog struct {
	age int
}

func (d dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog walk")
}

type pet1 struct {
	a    animal
	name string
}

type pet2 struct {
	animal
	name string
}
