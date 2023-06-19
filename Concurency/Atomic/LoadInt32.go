package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32 = 42

	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		defer wg.Done()

		value := atomic.LoadInt32(&counter)
		fmt.Println("Counter value:", value)

	}()

	go func() {
		defer wg.Done()

		value := atomic.LoadInt32(&counter)
		fmt.Println("Counter value:", value)

	}()
	go func() {
		defer wg.Done()

		value := atomic.LoadInt32(&counter)
		fmt.Println("Counter value:", value)

	}()
	go func() {
		defer wg.Done()

		value := atomic.LoadInt32(&counter)
		fmt.Println("Counter value:", value)

	}()

	wg.Wait()
}
