package main

import (
	"fmt"
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var ready = false

func main() {
	var w2 sync.WaitGroup
	w2.Add(5)
	go worker1(&w2)
	go worker2(&w2)
	go worker3(&w2)
	go worker4(&w2)
	go bose(&w2)
	w2.Wait()

}

func worker1(w2 *sync.WaitGroup) {
	defer w2.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Worker1 is working started")
}

func worker2(w2 *sync.WaitGroup) {
	defer w2.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Worker2 is working started")
}

func worker3(w2 *sync.WaitGroup) {
	defer w2.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Worker3 is working started")
}

func worker4(w2 *sync.WaitGroup) {
	defer w2.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Worker4 is working started")
}

func bose(w2 *sync.WaitGroup) {
	defer w2.Done()
	cond.L.Lock()
	defer cond.L.Unlock()
	time.Sleep(10 * time.Second)
	ready = true
	fmt.Println("Giveing instraction to workers.......")
	cond.Broadcast()
}
