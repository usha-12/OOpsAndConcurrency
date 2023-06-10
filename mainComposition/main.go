package main

import (
	"fmt"

	composition "example.com/package/Composition"
)

func main() {
	c := composition.NewCar("a", 600, 1000000)
	fmt.Println(c)
}
