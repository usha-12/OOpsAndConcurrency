package main

import "fmt"

func main() {
	//ch := make(chan type, capacity) how much data a channel can store before getting deadlock
	buffer := make(chan int, 3)
	buffer <- 10
	buffer <- 20
	buffer <- 30
	buffer <- 40
	fmt.Println(<-buffer)
	fmt.Println(<-buffer)
	fmt.Println(<-buffer)
	fmt.Println(<-buffer)

}
