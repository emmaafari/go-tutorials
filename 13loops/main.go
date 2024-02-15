package main

import "fmt"

func main() {
	fmt.Println("Welcome to the tutorial on loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednessday", "Thursday", "Friday"}

	fmt.Println(days)

	// 	for d := 0; d < len(days); d++ {
	// 		fmt.Println(days[d])
	// 	}

	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	for i, day := range days {
		fmt.Printf("Index is %v and value is %v\n", i, day)
	}

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 2 {
			goto message
		}
		rougueValue++
	}

message:
	fmt.Println("Hello from the goto message")
}
