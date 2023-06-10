package main

import (
	"fmt"
	"sync"
)

var count int = 0

func incrementCounter(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		count++
	}
	fmt.Println(count)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go incrementCounter(&wg)
	go incrementCounter(&wg)
	wg.Wait()

}
