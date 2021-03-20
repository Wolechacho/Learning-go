package main

func main() {
	//zero value of an interface - <nil>
	// var a animal
	// fmt.Println(a)

	//golang polymorphism
	// in := &indiaTax{
	// 	taxPercentage: 30,
	// 	income:        5000,
	// }
	// n := &nigeiraTax{
	// 	taxPercentage: 10,
	// 	income:        2000,
	// }

	// u := &usaTax{
	// 	taxPercentage: 10,
	// 	income:        3000,
	// }

	// taskSystems := []TaskSystem{in, n, u}
	// total := calculateTaxes(taskSystems)
	// fmt.Println("total : ", total)

	//custom types with interfaces
	// c := cat("smoley")
	// c.breathe()
	// c.walk()

	//multiple interfaces
	// var a animal
	// l := lion{}
	// a = l
	// a.breathe()
	// a.walk()
	// var m mammal
	// m = l
	// m.feed()

	//embedded interface in an interface
	// var h human

	// h = employee{name: "John"}
	// h.breathe()
	// h.walk()
	// h.speak()

	//embedding interface in struct
	// d := dog{age: 5}
	// p1 := pet1{name: "Milo", a: d}

	// fmt.Println(p1.name)
	// // p1.breathe()
	// // p1.walk()
	// p1.a.breathe()
	// p1.a.walk()

	// p2 := pet2{name: "Oscar", animal: d}
	// fmt.Println(p1.name)
	// p2.breathe()
	// p2.walk()
	// p1.a.breathe()
	// p1.a.walk()
}
