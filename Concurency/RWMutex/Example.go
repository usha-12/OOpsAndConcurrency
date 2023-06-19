package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data  int
	mutex sync.RWMutex
	wg    sync.WaitGroup
)

func main() {
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go readData(i)
	}
	wg.Add(1)
	go writeData()
	wg.Wait()
	fmt.Println("All goroutines finished execution ")

}
func readData(readerID int) {
	defer wg.Done()
	fmt.Printf("Reader %d is waiting to read \n", readerID)
	mutex.RLock() //data should we inconsistence
	fmt.Printf("Reader %d is reading the data: %d\n", readerID, data)
	time.Sleep(time.Second)
	mutex.RUnlock()
	fmt.Printf("Reader %d finished reading \n", readerID)
}

func writeData() {
	defer wg.Done()
	fmt.Println("Writer is waiting to write ")
	mutex.Lock()
	data = 42

	fmt.Println("Writer has written the data")

	time.Sleep(time.Second)
	mutex.Unlock()
	fmt.Println("writer finished writing")
}
