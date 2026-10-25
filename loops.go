package main

import "fmt"

func main() {
	for i := 1; i <= 30; i++ {

		if i == 23 {
			break

		}
		if i%5 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
