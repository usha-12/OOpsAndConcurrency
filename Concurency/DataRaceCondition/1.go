package main

import (
	"fmt"
	"time"
)

var counter int

func main() {
	go increment()
	go increment()
	go increment()
	go increment()
	go increment()

	time.Sleep(time.Second)

	fmt.Println("Counter value:", counter)
}

func increment() {
	for i := 0; i < 1000; i++ {
		counter++
	}
}
