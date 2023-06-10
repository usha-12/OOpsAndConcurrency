package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int) //created a channel
	go producer1(ch, &wg)
	go consumer2(ch, &wg)
	wg.Wait()

}

func producer1(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- 42
	fmt.Println("Send data")

}

func consumer2(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	value := <-ch // Receive a value from the channel

	fmt.Println("Received data", value)
}
