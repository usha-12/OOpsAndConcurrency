package main

func processSliceAndMap(numbers []int, data map[string]string) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	for key := range data {
		// Do something with the map values
	}

	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	data := make(map[string]string)
	data["a"] = "apple"
	data["b"] = "banana"
	total := processSliceAndMap(numbers, data)
	println("Total sum:", total)
}
