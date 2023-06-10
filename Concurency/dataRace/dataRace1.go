package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go increment1(&wg)
	go increment1(&wg)

	wg.Wait()
	fmt.Println(counter)

}

func increment1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		counter++
	}
	//fmt.Println(counter1)
}
