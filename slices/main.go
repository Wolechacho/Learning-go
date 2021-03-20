package main

func main() {
	//this throws an error
	//nilSlices()

	//this wont throw an error bcos you assign a pre-existing slice to it
	//acceptableSlice()

	//copies to noempty slices
	//nonEmptyCopySlices()

	//copies to empty slices does not populate the dest slice and the number of elements copied is zero
	//copytonilemptySlices()

	//clear the slice without clear the allocation in mem
	//clearSlicesWithoutAlloc()

	//clear with allocation
	//clearSlicesWitAlloc()

	//var x int
	//x,_ := exapmle2() //not work
	//x, _ = exapmle2() //work
	//x,y := exapmle2() //- work
	//x, y = exapmle2() //not work

	//fmt.Println(x)
	//fmt.Println(y)
}
