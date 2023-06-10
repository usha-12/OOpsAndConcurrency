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
	wg.Add(4)
	go emp1(&wg)
	go emp2(&wg)
	go emp3(&wg)
	go box(&wg)
	wg.Wait()

}

func emp1(wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Employee 1 ready")
}

func emp2(wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Empoyee 2 ready")
}
func emp3(wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Employee 3 ready")

}

func box(wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	time.Sleep(50 * time.Second)
	ready = true
	fmt.Println("Bose giving order to empoyees")
	cond.Broadcast()
}
