package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Golang!")
	time.Sleep(5 * time.Second)
	go printMessage()

}

func printMessage() {
	fmt.Println("Hello from goroutine!")
}
