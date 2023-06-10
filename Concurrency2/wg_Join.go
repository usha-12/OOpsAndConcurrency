package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		work1()

	}()
	wg.Wait()

	fmt.Println("Eclips:", now)
	fmt.Println("Printing main")

}

func work1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Printing work...............")
}
