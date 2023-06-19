package main

import (
	"fmt"
	"sync"
)

var count = 0

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg, &mutex)
	}
	wg.Wait()
	fmt.Println("count:", count)

}

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	count = count + 1
	mutex.Unlock()

}
