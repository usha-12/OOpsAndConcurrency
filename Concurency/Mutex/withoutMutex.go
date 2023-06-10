package main

import (
	"fmt"
	"sync"
)

var balance int = 500
var w6 sync.WaitGroup

func main() {

	w6.Add(3)
	go withdrow(100)
	go deposit(100)
	go checkBalance()

	//time.Sleep(2 * time.Second)

	w6.Wait()
	fmt.Println("Final balance:", balance)

}

func withdrow(x int) {
	defer w6.Done()
	balance = balance - x
	//fmt.Println("withdrow balance:", balance)

}

func deposit(x int) {
	defer w6.Done()
	balance = balance + x
	//fmt.Println("deposit balance:", balance)

}

func checkBalance() {
	defer w6.Done()
	fmt.Println("total balance:", balance)

}
