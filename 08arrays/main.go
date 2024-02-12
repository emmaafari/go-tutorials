package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitsList [4]string

	fruitsList[0] = "Apple"
	fruitsList[1] = "Mango"
	fruitsList[2] = "Tomato"
	fruitsList[3] = "Orange"

	fmt.Println("Fruit list: ", fruitsList)

	var vegetableList = [2]string{
		"Carrot",
		"Lettuce",
	}

	fmt.Println("Vegetables list: ", vegetableList)

}
