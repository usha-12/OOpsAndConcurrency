package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3) //conter of goroutine how much go routine finished
	go func1(&wg)
	go func2(&wg)
	go func3(&wg)
	wg.Wait()

}

func func1(wg *sync.WaitGroup) {
	defer wg.Done() //
	fmt.Println("Function 1 calling")

}

func func2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Function 2 calling")

}

func func3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Function 3 calling")

}
