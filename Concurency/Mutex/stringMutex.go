package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {
	var w5 sync.WaitGroup
	var str string
	w5.Add(2)

	go printHello(&str, &w5)
	go printWord(&str, &w5)

	w5.Wait()
	fmt.Println(str)

}

func printHello(str *string, w5 *sync.WaitGroup) {
	defer w5.Done()
	mutex.Lock()
	*str = *str + "Hello"
	mutex.Unlock()

}

func printWord(str *string, w5 *sync.WaitGroup) {
	defer w5.Done()
	mutex.Lock()
	*str = *str + "World"
	mutex.Unlock()

}
