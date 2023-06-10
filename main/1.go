package main

func processMap(data map[string]int) int {
	total := 0
	for _, value := range data {
		total += value
	}
	return total
}

func main() {
	data := make(map[string]int)
	data["a"] = 10
	data["b"] = 20
	data["c"] = 30
	result := processMap(data)
	println("Result:", result)
}
