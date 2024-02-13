package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Orange"}
	fmt.Println("Fruits lists 1: ", fruitList)

	fruitList = append(fruitList, "mango", "Banana")
	fmt.Println("Fruits lists 2: ", fruitList)

	//Get the item starting from index 1
	//Also the index after the colon is omitted i.e index 3 will not be included
	fruitList = append(fruitList[1:3])
	fmt.Println("Fruits lists 3 [1:3]: ", fruitList)

	highScore := make([]int, 4)
	highScore[0] = 237
	highScore[1] = 426
	highScore[2] = 574
	highScore[3] = 957
	// highScore[4] = 237 // will not work

	highScore = append(highScore, 500, 300, 600)
	// fmt.Println(highScore)

	sort.Ints(highScore)

	fmt.Println(highScore)
	// fmt.Println(sort.IntsAreSorted(highScore))

	//How to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "ruby", "python"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
