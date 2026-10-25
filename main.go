package main

import "fmt"

func main() {

	var a int
	var b int
	fmt.Println("Enter first number")
	fmt.Scan(&a)
	fmt.Println("Enter second number")
	fmt.Scan(&b)
	fmt.Println("Sum", a+b)
	fmt.Println("Differenc", a-b)
	fmt.Println("Product", a*b)
	if b != 0 {
		fmt.Println("Division", a/b)
	} else {
		fmt.Println("Cannot divison by zero")
	}
	if b != 0 {
		fmt.Println("Remainder", a%b)
	} else {
		fmt.Println("Cannot remaider by zero")
	}

}
