package main

import "fmt"

type cat string

type cattas interface {
	breathe()
	walk()
}

func (c cat) breathe() {
	fmt.Println("cat breathes")
}

func (c cat) walk() {
	fmt.Println("cat walks")
}
