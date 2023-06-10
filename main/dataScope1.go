package main

func processNumbers(numbers []int) int {
	sum := 0                      //stack
	for _, num := range numbers { //num = stack
		sum += num
	}
	return sum //stack
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}  //stack- refferncing underlying array- for slice data
	total := processNumbers(numbers) //stack- number is a slice types so inside of slice refferencing underlying array in heap area
	println("Total sum:", total)     //stack
}
