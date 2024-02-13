package main

import "fmt"

func main() {
	fmt.Println("Welcome to the tutorial on ifelse")

	loginCount := 23

	var result string
	if loginCount < 10 {
		result = "Regular ser"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println("Result: ", result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
}
