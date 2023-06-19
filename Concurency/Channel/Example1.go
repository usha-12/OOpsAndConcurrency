package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch1 := make(chan string)
	go func1(&wg, ch1)
	go func2(&wg, ch1)
	wg.Wait()

}

func func1(wg *sync.WaitGroup, ch1 chan<- string) {
	defer wg.Done()
	ch1 <- "Hello"
	fmt.Println("Sending data to funct2.....")

}

func func2(wg *sync.WaitGroup, ch1 <-chan string) {
	defer wg.Done()
	a := <-ch1
	fmt.Println("Recived data from func1:", a)

}
