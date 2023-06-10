package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})
	go task111(done)
	go task211(done)
	go task311(done)
	go task411(done)
	<-done
	<-done
	<-done
	<-done
	fmt.Println("total time:", time.Since(now))

}

func task111(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task1 complited")
	done <- struct{}{}
}

func task211(done chan struct{}) {

	fmt.Println("Task2 complited")
	done <- struct{}{}
}

func task311(done chan struct{}) {
	fmt.Println("Task3 complited")
	done <- struct{}{}
}

func task411(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task4 complited")
	done <- struct{}{}
}
