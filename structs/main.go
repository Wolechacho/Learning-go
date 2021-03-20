package main

func main() {
	//Methods on Anonymous nested struct fields
	// address := address1{city: "London", country: "UK"}

	// emp := employee1{name: "Sam", age: 31, salary: 2000, address1: address}
	// emp.details()

	// emp.address1.details()

	//chaining methods
	emp := employee1{name: "Sam", age: 31, salary: 2000}
	emp.printName().printAge().printSalary()
}
