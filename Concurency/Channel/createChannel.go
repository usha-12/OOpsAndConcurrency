package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go employee1(&wg)
	go employee2(&wg)
	go employee3(&wg)

	wg.Wait()

}

func employee1(wg *sync.WaitGroup) {
	defer wg.Done()
	var a int = 10
	fmt.Println("Print:", a)
	fmt.Println("Employee 1")

}

func employee2(wg *sync.WaitGroup) {
	defer wg.Done()
	var b int

	fmt.Println("Employee 2")

}

func employee3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Employee 3")

}
