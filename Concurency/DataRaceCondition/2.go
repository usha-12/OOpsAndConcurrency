package main

import (
	"fmt"
	"sync"
)

var counter1 int
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go increment1()
	go increment1()
	wg.Wait()

	fmt.Println("Final counter value:", counter1)
}

func increment1() {
	defer wg.Done()

	for i := 0; i < 100000; i++ {
		counter1++
	}
}
