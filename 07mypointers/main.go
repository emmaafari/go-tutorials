package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers.")

	myNumber := 100
	// & means reference
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is: ", *ptr)
	fmt.Println("New value is: ", myNumber)
}
