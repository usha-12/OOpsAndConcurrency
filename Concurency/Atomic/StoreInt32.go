package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		atomic.StoreInt32(&counter, 42)

	}()

	go func() {
		defer wg.Done()
		atomic.StoreInt32(&counter, 42)

	}()
	go func() {
		defer wg.Done()
		atomic.StoreInt32(&counter, 42)

	}()
	go func() {
		defer wg.Done()
		atomic.StoreInt32(&counter, 42)

	}()

	wg.Wait()

	fmt.Println("Counter value:", counter)

}
