package main

import "fmt"

func main() {

	numbers := [5]int{3, 6, 9, 12, 15}
	
	for index, value := range numbers{
		fmt.Printf("Индекс: %d, значение %\n" index ,value)
		
	}
}
