package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
}

type Address struct {
	city  string
	state string
}

type Person struct {
	name    string
	age     int
	address Address
}

func PrintEmployeeInfo() {
	emp := Employee{
		firstName: "Wole",
		lastName:  "Charles",
		age:       28,
	}

	fmt.Println(emp)
}

func PrintPersonInfo() {
	p := Person{
		name: "Naveen",
		age:  50,
		address: Address{
			city:  "Chicago",
			state: "Illinois",
		},
	}
	fmt.Println(p)
}
