## CHANNELS

Channels can be thought as pipes using which Goroutines communicate

THe main method is also a goroutine

When a data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel. Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.

Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel

just know nothing is executing below the write and read channel control except is closed
Channels are nil by default.The zero value of channel is nil

var ch chan int - these specifiies a nil chan
var ch chan int = make(chan int) - this does a declared channel

You cannot read or write into a nil channel - https://play.golang.org/p/4NO7qqS8xjR
you cannnot close a channel - https://play.golang.org/p/APFDpsA5rxp

Difference between closed and nil channel - https://nanxiao.gitbooks.io/golang-101-hacks/content/posts/nil-channel-vs-closed-channel.html

https://play.golang.org/p/DV7ZkcwA7SM

Wait groups are goroutines orchestrators

https://inconshreveable.com/07-08-2014/principles-of-designing-go-apis-with-channels/
https://notes.shichao.io/gopl/ch8/#channels

Lesson
DO NOT READ OR WRITE TO A CHANNEL WHETHER IT IS THE MAIN FUNCTION OR A CUSTOM FUNCTION YOU CREATE (BEFORE) YOU SPIN A GOROUTINE FUNCTION.please note before is the keyword

The unidirectional channel is used to provide the type-safety of the program so, that the program gives less error.unidirectional channels are just for compile checks

always make when you have multiple go routines,u should use WaitGroups and the Wait Done method should the called at the receiving goroutine function
always close the channel from the goroutine the sends value to the channel

Operation type Nil channel closed
Send Block Panic
Receive Block Not block,return zero value of channel type
Close panic panic

In Go you can also create a buffered channel. A buffered channel has some capacity to hold data hence for a buffered channel:

Send on a buffer channel only blocks if the buffer is full
Receive is only block when channel is empty

https://golangbyexample.com/all-operations-channel-golang/

You can do a send and receive operation in select in golang
