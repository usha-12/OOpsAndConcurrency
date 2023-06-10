package main

import (
	"fmt"
)

func main() {
	r := reverseString("hello")
	fmt.Println(r)

}

func reverseString(str string) string {

	newString := ""

	for i := len(str) - 1; i >= 0; i-- {
		newString = newString + string(str[i])
	}
	return newString

}
