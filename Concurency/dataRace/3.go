package main

import (
	"fmt"
	"sync"
)

var count = 0

func main() {
	var wg sync.WaitGroup
	wg.Add(6)
	go countIncrease(&wg)
	go countIncrease(&wg)
	go countIncrease(&wg)
	go countIncrease(&wg)
	go countIncrease(&wg)
	go countIncrease(&wg)
	wg.Wait()
	fmt.Println(count)

}

func countIncrease(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		count++
	}

	//fmt.Println(count)
}
