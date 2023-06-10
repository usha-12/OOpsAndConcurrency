package main

import (
	"fmt"
	"sync"
)

func main() { //why we are using waitgroup
	// we are using waitgroup to synchronize goroutines with our main function
	var wg sync.WaitGroup
	wg.Add(2)
	go print1(&wg)
	go print2(&wg)
	wg.Wait()
}

func print1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello1")
}
func print2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello2")
}
