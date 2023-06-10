package main

import (
	"encoding/json"
	"fmt"
)

type Persone struct {
	Name   string
	Age    int
	Gender rune
}

func (p Persone) dataRetunInString() string {
	return fmt.Sprintf("%s,%d,%s", p.Name, p.Age, string(p.Gender))

}

func (p Persone) marshellingData() string {
	jsonData, _ := json.Marshal(p)
	fmt.Println("Json Data:", string(jsonData))
	return string(jsonData)

}
func main() {
	person1 := Persone{Name: "Usha", Age: 23, Gender: 'F'}
	func1 := person1.dataRetunInString()
	fmt.Println(func1)
	func2 := person1.marshellingData()
	fmt.Println(func2)
}
