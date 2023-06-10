package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	go task11()
	go task21()
	go task31()
	go task41()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("total time:", time.Since(now))

}

func task11() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task1 complited")
}

func task21() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task2 complited")
}

func task31() {
	fmt.Println("Task3 complited")
}

func task41() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task4 complited")
}
