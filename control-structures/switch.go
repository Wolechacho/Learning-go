package main

import (
	"fmt"
	//"mygoplayground/folderfunction"
)

//Switch - learning switch statement
func Switch() {
	a := 4

	switch a {
	case 1:
		fmt.Println("a is 1 (Switch)")
	case 2:
		fmt.Println("a is 2 (Switch)")
	case 3:
		fmt.Println("a is 3 (Switch)")
	case 4:
		fmt.Println("a is 4 (Switch)")
	}
}

//DefaultSwitch - learning the use of default case with switch statement
func DefaultSwitch() {
	a := 5

	switch a {
	case 1:
		fmt.Println("a is 1 (Default Switch)")
	case 2:
		fmt.Println("a is 2 (Default Switch)")
	case 3:
		fmt.Println("a is 3 (Default Switch)")
	case 4:
		fmt.Println("a is 4 (Default Switch)")
	default:
		fmt.Println("Value not found (Default Switch)")
	}
}

//DefaultSwitchWithNoCondition - learning switch statement with no condition
func DefaultSwitchWithNoCondition() {
	a := 3
	switch {
	case a == 1:
		fmt.Println("a is 1 (switch with no condition)")
	case a == 2:
		fmt.Println("a is 2 (switch with no condition)")
	case a == 3:
		fmt.Println("a is 3 (switch with no condition)")
	case a == 4:
		fmt.Println("a is 4 (switch with no condition)")
	}
}

//DefaultSwitchWithInitializedVar - learning switch statement with initialized variable
func DefaultSwitchWithInitializedVar() {

	switch a := 2; {
	case a == 1:
		fmt.Println("a is 1 (switch with initialized variable)")
	case a == 2:
		fmt.Println("a is 2 (switch with initialized variable)")
	case a == 3:
		fmt.Println("a is 3 (switch with initialized variable)")
	case a == 4:
		fmt.Println("a is 4 (switch with initialized variable)")
	}
}

//MultipleExpression per case
func MultipleExpression() {
	weekDay := "tuesday"

	switch weekDay {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("working day")
	default:
		fmt.Println("Not working day: weekend")
	}

}

//BreakStmtSwitch : Break statement with switch
//Real world application - check the occurrence of a value just one time
func BreakStmtSwitch() {
	var numbArray = [3]int{1, 2, 5}
MyLoopName:
	for _, numb := range numbArray {
		switch numb {
		case 1:
			fmt.Println("1 matched")
			break MyLoopName
		case 2:
			fmt.Println("2 matched")
		case 0:
			fmt.Println("0 matched")
		}
	}
}

//SwitchFxnExpression - switch case with function
// func SwitchFxnExpression() {
// 	switch folderfunction.Myfunction("mon") {
// 	case folderfunction.Myfunction("mon"), folderfunction.Myfunction("tue"), folderfunction.Myfunction("wed"), folderfunction.Myfunction("thu"), folderfunction.Myfunction("fri"):
// 		fmt.Println("working day")
// 	case folderfunction.Myfunction("sat"), folderfunction.Myfunction("sun"):
// 		fmt.Println("Not working day: weekend")
// 	}
// }

//TypeAssertion - for type assertions
func TypeAssertion() {
	var interfaceDemo interface{} = "true"

	switch interfaceDemo.(type) {
	case string:
		fmt.Println("string type matched")
	case int:
		fmt.Println("int type matched")
	default:
		fmt.Println("default type matched")

	}
}

//TypeAssertionWithStruct - with struct
func TypeAssertionWithStruct() {
	type Employee struct {
		ID   int
		Name string
	}
	emp := Employee{1, "Kiran"}
	var empInterface interface{}
	empInterface = emp
	switch empInterface.(type) {
	case Employee:
		fmt.Println("Employee type matched")
	case string:
		fmt.Println("String type matched")
	default:
		fmt.Println("default type matched")
	}
}
