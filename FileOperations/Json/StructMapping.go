package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name string `json:"name"`
	Age  string `json:"ageofperson"`
}

func main() {
	jsonStr := `{"name":"Usha","ageofperson":"23"}`

	var data Data

	json.Unmarshal([]byte(jsonStr), &data)
	fmt.Printf("Name %s, age %s:", data.Name, data.Age)
}
