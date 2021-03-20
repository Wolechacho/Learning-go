package main

import "fmt"

type employee1 struct {
	name   string
	age    int
	salary int
	address1
}

type address1 struct {
	city    string
	country string
}

func (a address1) details() {
	fmt.Printf("City: %s\n", a.city)
	fmt.Printf("Country: %s\n", a.country)
}

func (e employee1) printName() employee1 {
	fmt.Printf("Name: %s\n", e.name)
	return e
}

func (e employee1) printAge() employee1 {
	fmt.Printf("Age: %d\n", e.age)
	return e
}

func (e employee1) printSalary() {
	fmt.Printf("Salary: %d\n", e.salary)
}
