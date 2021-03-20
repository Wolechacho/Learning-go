package main

var (
	ch2     chan int    = make(chan int)
	mychanl chan string = make(chan string)
)

func main() {
	//Print a nil channel
	//printNilChannel()

	//print a initialized channel
	//printInitializedChannel()

	//more on channels

	//this is going to throw an error
	//intro()

	//correct way
	// go intro2()
	// fmt.Println(<-ch2)

	//another one
	//PrintChannelMessage()

	//how to check if the channel is read and done
	// b := make(chan bool)
	// go Hello(b)
	// b <- true
	// fmt.Println("Values printed")

	//BUFFERED CHANNEL
	//bufferChannel()

	//blocked buffered channel
	//blockedBufferChannel()

	//select with channels and goroutines
	//execute()

	///this select statement blocks
	//blockSelect()

	//sendOperationOnSelect()
	//fmt.Println("Done")

	//select with default case
	//selectWithDefaultCase()

	//select with default case with for loop
	// ch12 := make(chan string)
	// go goThree(ch12)
	// selectWithForLoopDefaultCase(ch12)

	//select with block timeout
	SelectWithBlockTimeout()
}
