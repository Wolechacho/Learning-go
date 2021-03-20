package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func execute() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
	fmt.Println("Done")
}

func goOne(ch chan string) {
	ch <- "From goOne goroutine"
}

func goTwo(ch chan string) {
	msg := <-ch
	fmt.Println(msg)
}

func blockSelect() {
	ch1 := make(chan string)
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	}
}

//sending operation and also receiving on a select
func sendOperationOnSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go goOne(ch1)
	go goTwo(ch2)
	select {

	case msg1 := <-ch1:
		fmt.Println(msg1)
	case ch2 <- "To goTwo goroutine":
	}
}

//default case with select
func selectWithDefaultCase() {
	ch1 := make(chan string)
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	default:
		fmt.Println("Default statement executed")
	}
}

//default case with select with For loop
//note always use the return statement in the case statement
func selectWithForLoopDefaultCase(s chan string) {
	for {
		select {
		case msg := <-s:
			fmt.Println(msg)
			return
		default:
			fmt.Println("Default statement executed")
		}
	}

}

func goThree(ch1 chan string) {
	ch1 <- "Helo there"
}

//Select with blocking timeout
func SelectWithBlockTimeout() {
	ch1 := make(chan string)
	go goFour(ch1)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout")
	}
}

func goFour(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "From goFour goroutine"
}
