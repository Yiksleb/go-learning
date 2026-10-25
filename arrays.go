package main

import "fmt"

func main() {

	numbers := [3][2]int{
		{5, 4},
		{3, 5},
		{4, 4},
	}
	for index, row := range numbers {
		sum := 0
		for _, value := range row {
			sum += value

		}
		fmt.Println("Sum of row", index, ":", sum)
	}
}
