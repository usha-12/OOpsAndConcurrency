package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var cond = sync.NewCond(&sync.Mutex{})
	var ready = false
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		cond.L.Lock()
		defer cond.L.Unlock()
		for !ready {
			cond.Wait()

		}
		fmt.Println("Goroutine 1: Ready signal received")

	}()

	go func() {
		defer wg.Done()

		cond.L.Lock()
		defer cond.L.Unlock()

		time.Sleep(50 * time.Second)
		ready = true

		fmt.Println("Goroutine 2: Ready signal sent")
		cond.Broadcast()

	}()
	wg.Wait()

}
