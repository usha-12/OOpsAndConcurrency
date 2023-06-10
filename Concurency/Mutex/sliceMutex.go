package main

import (
	"fmt"
	"sync"
)

func main() {
	var w1 sync.WaitGroup
	var mutex sync.Mutex
	slice := []string{}
	w1.Add(2)
	go func(w1 *sync.WaitGroup, mux *sync.Mutex) {
		defer w1.Done()

		mux.Lock()
		defer mux.Unlock()
		slice = addHello(slice)

	}(&w1, &mutex)
	go func(w1 *sync.WaitGroup, mux *sync.Mutex) {
		defer w1.Done()
		mux.Lock()
		defer mux.Unlock()
		slice = addWorld(slice)

	}(&w1, &mutex)

	fmt.Println(slice)
	w1.Wait()

}
func addHello(slice1 []string) []string {
	slice1 = append(slice1, "Hello")
	return slice1

}

func addWorld(slice2 []string) []string {
	slice2 = append(slice2, "World")
	return slice2
}
