package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.co"}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()

}

func getStatusCode(endpoint string) {

	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOps in endpoint")

	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
