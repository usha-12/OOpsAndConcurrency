package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go work3(wg)
	wg.Wait()

}

func work3(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Work is done")

}
