package main

import (
	"fmt"
	"sync"
)

type Data1 struct {
	Value int
}

func main() {
	data := Data{}
	var w4 sync.WaitGroup
	w4.Add(1)

}
func producer3(w4 *sync.WaitGroup) {
	defer w4.Done()
	for i := 0; i < 100; i++ {
		fmt.Println("Producer sending", i)

	}

}

func consumer3(w4 *sync.WaitGroup) {
	defer w4.Done()

}
