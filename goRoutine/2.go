package main

import (
	"fmt"
	"time"
)

func main() {
	go say("Hello", 3)
	go say("Hello2", 2)
	fmt.Scanln()

}

func say(s string, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(i, s)
	}
}
