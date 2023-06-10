package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

func main() { //why we are using waitgroup
	// we are using waitgroup to synchronize goroutines with our main function

	wg1.Add(2)
	go print11()
	go print21()
	wg1.Wait()

}

func print11() {
	defer wg1.Done()
	fmt.Println("hello1")
}
func print21() {
	defer wg1.Done()
	fmt.Println("Hello2")
}
