package main

import (
	"fmt"
	"time"
)

func main() {
	go f1()
	go f2()
	fmt.Println("Hello world")
	time.Sleep(1 * time.Second) //explicitly

}

func f1() {
	fmt.Println("Hello")

}

func f2() {
	fmt.Println("HIII!")
}
