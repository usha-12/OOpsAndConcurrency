package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go printHello(&wg)
	go printWorld(&wg)
	go printHelloWorld(&wg)
	wg.Wait()

}

func printHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Print hello")

}

func printWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Print World")

}

func printHelloWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Print hello world")

}
