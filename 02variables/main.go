package main

import "fmt"

func main() {
	var username string = "Emmanuel Kofi"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var smallFloat float32 = 255.466288564535654
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
}
