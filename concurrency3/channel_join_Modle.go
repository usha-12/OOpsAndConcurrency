package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})

	go func() {
		work3()
		done <- struct{}{}

	}()
	<-done
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("Done waiting,main exits")

}

func work3() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing some stuff")
}
