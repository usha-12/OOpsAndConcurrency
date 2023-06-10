package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var count int

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumber(&wg)
	go printNumber1(&wg)
	wg.Wait()
	fmt.Println(count)

}

func printNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		count++
		mutex.Unlock()

	}

}

func printNumber1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i < 20; i++ {
		mutex.Lock()
		count++
		mutex.Unlock()
	}
}
