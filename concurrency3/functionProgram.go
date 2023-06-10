package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	task1()
	task2()
	task3()
	task4()

	fmt.Println("total time:", time.Since(now))

}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task1 complited")
}

func task2() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task2 complited")
}

func task3() {
	fmt.Println("Task3 complited")
}

func task4() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task4 complited")
}
