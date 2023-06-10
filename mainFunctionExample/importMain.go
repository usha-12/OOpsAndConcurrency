package main

import (
	"fmt"

	Struct1 "example.com/package/import"
	//_ "example.com/package/import"
)

func main() {
	//p := Struct1.Person{}
	p := Struct1.NewPerson("shivani", "Ahirwar", 23)

	//p.SetFristName("Usha")
	fmt.Println("First Name is :", p.GetFirstName())

}
