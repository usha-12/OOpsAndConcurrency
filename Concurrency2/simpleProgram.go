package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	go task1()
	go task2()
	go task3()
	go task4()
	go task5()
	time.Sleep(time.Second)
	fmt.Println(time.Since(now))

}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task1")
}
func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Task2")
}
func task3() {
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Task3")
}
func task4() {
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Task4")
}
func task5() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Task5")
}
