package main

import "fmt"

func main() {

	go printNumber(10)

}

func printNumber(number int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
