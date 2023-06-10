package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var myArray [500]string

	for i := 0; i < len(myArray); i++ {
		myArray[i] = randomString(10) // Generate a random string of length 10
	}

	for i := 0; i < len(myArray); i++ {
		fmt.Println("String without revese:", myArray[i])
		//reverseString := revese()
		//fmt.Println("reverse string:", reverseString)
		go revese(myArray[i])
	}

}

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	//fmt.Println(string(b))
	return string(b)
}

func revese(str string) {
	newString := ""

	for i := len(str) - 1; i >= 0; i-- {
		newString = newString + string(str[i])
	}
	//return newString
	fmt.Println(newString)

}
