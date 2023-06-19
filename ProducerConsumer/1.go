package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	cha1 := make(chan float64)
	wg.Add(2)
	go reciver(&wg, cha1)
	go consumer(&wg, cha1)
	wg.Wait()

}

func reciver(wg *sync.WaitGroup, cha1 chan<- float64) {
	defer wg.Done()
	cha1 <- 19.90
	fmt.Println("Sending data......")

}

func consumer(wg *sync.WaitGroup, cha1 <-chan float64) {
	defer wg.Done()
	data := <-cha1
	fmt.Println("Received data from reciver:", data)

}
