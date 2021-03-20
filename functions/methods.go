package main

import "fmt"

type Bike struct {
	Model string
	Size  int
}

func (b Bike) ValueReceiver() {
	fmt.Printf("VALUE :: The address of the received bike is: %p\n", &b)
	// Address of the object is different than the ones in main
	// Changing the object inside the scope here won't reflect to the caller
	b.Model = "BMW"
	fmt.Println("Inside ValueReceiver model: ", b.Model)
}

func (b *Bike) PointerReceiver() {
	fmt.Printf("POINTER :: The address of the received bike is: %p\n", b)
	// Address of the object is the same as the ones in main
	b.Model = "BMW"
	fmt.Println("Inside PointerReceiver model: ", b.Model)
}

func PrintBike() {
	v := Bike{"Honda CBR", 650}
	p := &Bike{"Suzuki V-Storm", 650}

	fmt.Printf("Value object address in main: %p\nPointer object address in main: %p\n\n", &v, p)
	v.ValueReceiver()
	fmt.Println("Value model outside the function: ", v.Model)

	fmt.Println("")
	p.PointerReceiver()
	fmt.Println("Pointer model outside the function: ", p.Model)
}

//change the state of an object
//memory management
