package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32 = 0

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		swapped := atomic.CompareAndSwapInt32(&counter, 0, 10)
		if swapped {
			fmt.Println("Counter was succussfully swapped to 10")

		} else {
			fmt.Println("Counter was not swapped")
		}

	}()

	go func() {
		defer wg.Done()
		swapped := atomic.CompareAndSwapInt32(&counter, 0, 20)
		if swapped {
			fmt.Println("Counter was successfully swapped to 20")

		} else {
			fmt.Println("Counter was not swapped")
		}

	}()

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
