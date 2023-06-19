package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	cha := make(chan int, 3)
	go function1(&wg, cha)
	go function2(&wg, cha)
	wg.Wait()
}

func function1(wg *sync.WaitGroup, cha chan<- int) {
	defer wg.Done()
	cha <- 10000
	cha <- 20000
	cha <- 30000
	fmt.Println("Sending data")

}

func function2(wg *sync.WaitGroup, cha <-chan int) {
	defer wg.Done()
	a := <-cha
	//b := <-cha
	//c := <-cha
	fmt.Println("Recived data:", a)
	fmt.Println("Recived data:", a)
	fmt.Println("Recived data:", a)
}
