package main

import "fmt"

func main() {
	fmt.Println("Welcome to the tutorial on defer")
	//Defer will cause the
	defer fmt.Println("World")
	fmt.Println("Hello")

	defer myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		fmt.Println("Value is:", i)
	}
}
