package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan string)
	go func1(&wg, ch)
	go func2(&wg, ch)
	wg.Wait()

}

func func1(wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	ch <- "Hello"
	fmt.Println("Sending Data from Function2....")

}

func func2(wg *sync.WaitGroup, ch <-chan string) {
	defer wg.Done()
	a := <-ch
	fmt.Println("Recceived data from function1:", a)

}
