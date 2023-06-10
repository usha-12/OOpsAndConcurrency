package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	book := Book{
		Title:  "5 AM clab",
		Name:   "Early win",
		Author: "Robin Sharma",
	}
	fmt.Printf("%+v", book)

	jsonData, _ := json.Marshal(book)
	fmt.Println(string(jsonData))

	var newBook Book

	err := json.Unmarshal(jsonData, &newBook)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("New Book:", newBook)

}

type Book struct {
	Title  string
	Name   string
	Author string
}
