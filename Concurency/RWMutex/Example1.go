package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data1  int
	mutex1 sync.RWMutex
	wg1    sync.WaitGroup
)

func main() {
	//launch 3 reader goroutines

	for i := 1; i <= 3; i++ {
		wg1.Add(1)
		go readData1(i)
	}

	//launch 2 writer goroutines

	for i := 1; i <= 2; i++ {
		wg1.Add(1)
		go writeData1(i)
	}

	wg1.Wait()
	fmt.Println("All goroutines finished execution..")
}

func readData1(readerID int) {
	defer wg1.Done()
	fmt.Printf("Reader %d is waiting to read \n", readerID)
	mutex1.RLock() //acquire a read lock

	fmt.Printf("Reader %d is reading the data %d\n", readerID, data1)

	time.Sleep(time.Second) //simulate some work

	mutex1.RUnlock() //release the read lock

	fmt.Printf("Ewder %d finished reading\n", readerID)

}

func writeData1(writerID int) {
	defer wg1.Done()

	fmt.Printf("Writer %d is waiting to write\n", writerID)
	mutex1.Lock() //acquire a write lock
	data1 = writerID
	fmt.Printf("Writer %d has written the data \n", writerID)

	time.Sleep(time.Second) //simulate some work

	mutex1.Unlock() //release the write lock

	fmt.Printf("Writer %d finished writing \n", writerID)

}
