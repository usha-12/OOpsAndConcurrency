package main

import (
	"fmt"
	"time"
)

func main() {
	go printMessage2()
	go printMessage2()
	go printMessage2()
	go printMessage2()
	time.Sleep(10 * time.Millisecond)

}

func printMessage2() {
	fmt.Println("Hello!")
}
