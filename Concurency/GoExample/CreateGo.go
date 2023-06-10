package main

import (
	"fmt"
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var ready = false

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go printMessage1(&wg)
	go printMessage2(&wg)
	go printMessage3(&wg)
	wg.Wait()

}

func printMessage1(wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Print emp1...")
}

func printMessage2(wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Print emp2..")
}

func printMessage3(wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	time.Sleep(50 * time.Second)
	ready = true
	fmt.Println("bose giving order to emp1 and emp2..")
	cond.Broadcast()

}
