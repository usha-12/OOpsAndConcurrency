package main

import (
	"fmt"
	"time"
)

func main() {

	go work() //fork point
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done wating main exits.......")
	//join point

}

func work() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Printing work...............")
}
