package main

import (
	"fmt"
	"time"
)

func main() {
	go printPrint()
	go func() {
		fmt.Println("Hello this is anynomys goroutine")

	}()

	time.Sleep(time.Millisecond)

}

func printPrint() {
	fmt.Println("Hello print mesage")
}
