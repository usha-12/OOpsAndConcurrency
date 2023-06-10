package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printLatter(&wg)
	go printNumber(&wg)
	wg.Wait()

}

func printNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)

	}

}

func printLatter(wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 'a'; j <= 'z'; j++ {
		fmt.Printf("%c", j)
	}

}
