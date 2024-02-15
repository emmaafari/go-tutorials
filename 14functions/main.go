package main

import "fmt"

func main() {
	fmt.Println("Welcome to the tutorial on functions")
	greet()

	result := add(2, 3)
	fmt.Printf("Result is: %v \n", result)

	proResult := proAdd(1, 2, 3, 5)
	fmt.Printf("Pro result is: %v", proResult)
}

func add(x int, y int) int {
	return x + y
}

func proAdd(values ...int) int {
	total := 0

	for _, v := range values {
		total += v
	}

	return total
}

func greet() {
	fmt.Println("Hello from golang")
}
