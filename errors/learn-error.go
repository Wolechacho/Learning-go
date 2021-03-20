package main

import (
	"errors"
	"fmt"
	"os"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

//using error.New()
func sampError() {
	er := errors.New("error occured")
	fmt.Println(er)
}

//using fmt Errorf
func formatError() {
	sampleErr := fmt.Errorf("Err is: %s", "database connection issue")
	fmt.Println(sampleErr)
}

//creating custom error
type inputError struct {
	message      string
	missingField string
}

//return message
func (i *inputError) Error() string {
	return i.message
}

//return missingField
func (i *inputError) getMissingField() string {
	return i.missingField
}

//note returned error is an interface
func validate(name, gender string) error {
	if name == "" {
		return &inputError{message: "name is missing", missingField: "name"}
	}
	if gender == "" {
		return &inputError{message: "gender is missing", missingField: "gender"}
	}
	return nil
}

//ignore errors
func openfile() {
	file, _ := os.Open("non-existing.txt")
	fmt.Println(file)
}

//wrapping errors
type errorOne struct{}

func (e errorOne) Error() string {
	return "Error One happended"
}

//wrap errors
func wrapError() {
	e1 := errorOne{}

	e2 := fmt.Errorf("E2: %w", e1)

	e3 := fmt.Errorf("E3: %w", e2)

	fmt.Println(e2)

	fmt.Println(e3)
}

//unwrap errors
func unwrapError() {
	e1 := errorOne{}
	e2 := fmt.Errorf("E2: %w", e1)
	e3 := fmt.Errorf("E3: %w", e2)
	fmt.Println(errors.Unwrap(e3))
	fmt.Println(errors.Unwrap(e2))
	fmt.Println(errors.Unwrap(e1))
}

//check if two errors are equal
func do() error {
	return errorOne{}
}

//why you should use errors.Is for comparison
func do1() error {
	return fmt.Errorf("E2: %w", errorOne{})
}

//wrapping error 2

type notPositive struct {
	num int
}

func (e notPositive) Error() string {
	return fmt.Sprintf("checkPositive: Given number %d is not a positive number", e.num)
}

type notEven struct {
	num int
}

func (e notEven) Error() string {
	return fmt.Sprintf("checkEven: Given number %d is not an even number", e.num)
}

func checkPositive(num int) error {
	if num < 0 {
		return notPositive{num: num}
	}
	return nil
}

func checkEven(num int) error {
	if num%2 == 1 {
		return notEven{num: num}
	}
	return nil
}

func checkPostiveAndEven(num int) error {
	if num > 100 {
		return fmt.Errorf("checkPostiveAndEven: Number %d is greater than 100", num)
	}

	err := checkPositive(num)
	if err != nil {
		return fmt.Errorf("checkPostiveAndEven: %w", err)
	}

	err = checkEven(num)
	if err != nil {
		return fmt.Errorf("checkPostiveAndEven: %w", err)
	}

	return nil
}
