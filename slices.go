package main

import "fmt"

func main() {

	numbers := []int{12, 5, 8, 12, 3, 9, 9}
	numbers = removeNegatives(numbers)
	fmt.Println(numbers)

}

func removeNegatives(numbers []int) []int {
	result := []int{}

	for _, value := range numbers {
		if value >= 0 {
			result = append(result, value)

		}
	}
	return result
}
