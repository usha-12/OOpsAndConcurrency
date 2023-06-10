package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		work1()

	}()
	wg.Wait()

	fmt.Println("Done waiting,main exits")

}

func work1() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing some stuff")
}
