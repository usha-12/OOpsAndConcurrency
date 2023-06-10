package main

import (
	polymorphism "example.com/package/polimorphism"
)

func main() {
	var c polymorphism.Shape = polymorphism.Circle{}
	var s polymorphism.Shape = polymorphism.Square{}

	c.Render()
	s.Render()
}
