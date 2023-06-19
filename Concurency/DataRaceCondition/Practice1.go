package main

import (
	"fmt"
	"sync"
)

var count10 = 0

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go printMutex(&wg)
	go printMutex(&wg)
	go printMutex(&wg)
	go printMutex(&wg)
	go printMutex(&wg)
	wg.Wait()
	fmt.Println(count10)

}

func printMutex(wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		count10++

	}
	//fmt.Println(count10)
	defer wg.Done()

}

//chan2 chan<-int
// chan2 <- chan int
