package main

import (
	"fmt"
	"sync"
)

var (
	initialized bool
	value       int
	once        sync.Once
	wg          sync.WaitGroup
)

func main() {
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go accessValue(i)
	}

	wg.Wait()
	fmt.Println("all goroutines finished execution")

}

func accessValue(goroutineID int) {
	defer wg.Done()
	once.Do(initialize)
	//initialize()
	fmt.Printf("Goroutine %d: value is %d\n", goroutineID, value)
}

func initialize() {
	value = 42
	fmt.Println("Initialization complete")
}
