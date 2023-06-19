package main

import (
	"fmt"
	"sync"
)

func main() {
	cha := make(chan int)
	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		wg.Add(1)
		for i := 1; i <= 5; i++ {
			cha <- i
			//close(cha) //clouse the channel after sending all values

		}
		close(cha) //clouse the channel after sending all values

	}()

	for value := range cha {
		fmt.Println(value)
	}
	wg.Wait()

}
