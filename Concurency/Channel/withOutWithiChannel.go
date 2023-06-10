package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	Value int
}

func producer(data *Data, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("Producer sending:", i)
		data.Value = i          // Update shared data
		time.Sleep(time.Second) // Simulate some work
	}
}

func consumer(data *Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)                       // Simulate some work
		fmt.Println("Consumer received:", data.Value) // Access shared data
	}
}

func consumer1(data *Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Consumer1 received:", data.Value)
	}
}

func main() {
	data := &Data{} // Shared data
	var wg sync.WaitGroup

	wg.Add(3)
	go producer(data, &wg)
	go consumer(data, &wg)
	go consumer1(data, &wg)

	wg.Wait()
	fmt.Println("All done!")
}
