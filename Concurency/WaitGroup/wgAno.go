package main

import (
	"fmt"
	"sync"
)

func main() {
	var w sync.WaitGroup
	w.Add(1)
	go func(w *sync.WaitGroup) {
		defer w.Done()
		print111()

	}(&w)

	go func(w *sync.WaitGroup) {
		defer w.Done()

		print211()

	}(&w)

	w.Wait()

}

func print111() {

	fmt.Println("hello1")
}
func print211() {

	fmt.Println("Hello2")
}
