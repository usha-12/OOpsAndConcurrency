package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	cha := make(chan int)
	wg.Add(2)
	go firstFunction(&wg, cha)
	go secondFunction(&wg, cha)
	wg.Wait()

}

func firstFunction(wg *sync.WaitGroup, ch1 chan<- int) {
	defer wg.Done()
	ch1 <- 1000
	fmt.Println("Sending data to function....")

}

func secondFunction(wg *sync.WaitGroup, ch1 <-chan int) {
	defer wg.Done()
	data := <-ch1

	fmt.Println("Receiving data from function1:", data)

}
