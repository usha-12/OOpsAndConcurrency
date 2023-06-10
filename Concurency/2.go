package main

import (
	"fmt"
	"time"
)

func main() {
	go printMessage1("Hello")
	go printMessage1("World")
	time.Sleep(1 * time.Second)
}

func printMessage1(message string) {
	fmt.Println(message)
}
