package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg3 sync.WaitGroup
	wg3.Add(2)
	cha1 := make(chan float64)
	go printMsg1(cha1, &wg3)
	go printMsg2(cha1, &wg3)
	wg3.Wait()

}

func printMsg1(cha1 chan<- float64, wg3 *sync.WaitGroup) {
	defer wg3.Done()
	cha1 <- 12.89
	fmt.Println("Sending data .....")

}

func printMsg2(cha1 <-chan float64, wg3 *sync.WaitGroup) {
	defer wg3.Done()
	x := <-cha1
	fmt.Println("Recived data from function printMsg1:", x)

}
