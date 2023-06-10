package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go printMessage3(ch)
	a := <-ch
	fmt.Println(a)
	go printMessage4()
	//time.Sleep(10 * time.Second)

}

func printMessage3(ch chan string) {
	ch <- "Hello"
}
func printMessage4() {
	fmt.Println("Hello2")
}
