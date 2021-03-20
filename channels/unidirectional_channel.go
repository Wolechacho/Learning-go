package main

import (
	"fmt"
	"sync"
)

//UNIDIRECTIONAL CHANNEL

//SEND ONLY CHANNEL
//func main1() {
// var wg sync.WaitGroup
// wg.Add(1)
// go sending(mychanl)
// go receiving(mychanl,&wg)

// wg.Wait()
// fmt.Println("Done")


//ch := make(chan int,3)
//process(ch)
//fmt.Println(<-ch)

// ch1 := make(chan string, 1)
// ch1 <- "word"
// select {
// case msg := <-ch1:
// 	fmt.Println(msg)
// default:
// 	fmt.Println("Default statement executed")
// }
//}

func process(ch chan<- int) {
	ch <- 2
}

func sending(s chan<- string) {
	s <- "GeeksforGeeks"
	close(mychanl)
}

func receiving(s <-chan string, wg *sync.WaitGroup) {
	fmt.Println(<-s)
	wg.Done()
}
