package polymorphism

import "fmt"

type Shape interface {
	Render()
}

type Circle struct{}

func (C Circle) Render() {
	fmt.Println("Circle is getting rendered")
}

type Square struct{}

func (C Square) Render() {
	fmt.Println("Square is getting rendered")

}
