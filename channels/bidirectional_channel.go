package main

import "fmt"

//blocked buffered channel
func bufferChannel() {
	ch := make(chan int, 1)
	ch <- 1
	//ch <- 2
	fmt.Println("Sending value to channnel complete")
	val := <-ch
	fmt.Printf("Receiving Value from channel finished. Value received: %d\n", val)
}

func blockedBufferChannel() {
	ch := make(chan int, 1)
	ch <- 1
	ch <- 2
	fmt.Println("Sending value to channnel complete")
	val := <-ch
	fmt.Printf("Receiving Value from channel finished. Value received: %d\n", val)
}
