package main

import (
	"fmt"
)

func printNilChannel() {
	var a chan int
	fmt.Println(a)
}

func printInitializedChannel() {
	a := make(chan int)
	fmt.Println(a)
}

func intro() {
	ch := make(chan int)

	ch <- 3
	fmt.Println(<-ch)
}

func intro2() {
	ch2 <- 60
}

//PrintChannelMessage print channel message
func PrintChannelMessage() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	message := <-messages
	fmt.Println(message)
}

func Hello(done chan bool) {
	fmt.Println("Hello from hello function")
	_, ok := <-done

	if ok {
		fmt.Println("Done")
	}
}
