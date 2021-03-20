package main

import (
	"fmt"
	"sync"
)

var (
	x = 0
)
var ma = make(map[int]int)

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	// ch := make(chan bool, 1)
	// for i := 0; i < 1000; i++ {
	// 	w.Add(1)
	// 	//go increment(&w,&m)
	// 	go increment2(&w, ch)
	// }

	//fmt.Println("final value of x", x)

	w.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer w.Done()
			m.Lock()
			ma[i] = i
			fmt.Println(i,ma)
			m.Unlock()

		}()
	}

	w.Wait()
	fmt.Println(len(ma))
}

func increment(w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	//x = x + 1 is the critical section of the code
	x = x + 1
	m.Unlock()
	w.Done()
}

//using channels
func increment2(w *sync.WaitGroup, c chan bool) {
	c <- true
	//x = x + 1 is the critical section of the code
	x = x + 1
	<-c
	w.Done()
}
