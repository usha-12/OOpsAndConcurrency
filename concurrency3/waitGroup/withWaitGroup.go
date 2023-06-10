package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Numbers of cores:", runtime.NumCPU())
	wg.Add(10)
	now := time.Now()

	for i := 0; i < 10; i++ {
		go work1(&wg, i+1)
	}
	wg.Wait()
	fmt.Println("Total tine:", time.Since(now))
	fmt.Println("main is done")

}

func work1(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task", id, "Done")
}
