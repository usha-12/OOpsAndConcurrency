package main

import (
	"fmt"

	comparable "example.com/package/Comparable"
)

func main() {
	p1 := comparable.NewPerson("Usha", "Ahirwar", 23, []string{"alis"})
	p2 := comparable.NewPerson("Usha", "Ahirwar", 23, []string{"alisa"})

	result := p1.Equals(p2)
	if result {
		fmt.Println("Both the struct same")

	} else {
		fmt.Println("Not same")
	}

}
