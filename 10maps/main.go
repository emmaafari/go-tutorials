package main

import "fmt"

func main() {
	fmt.Println("Welcome to a tutorial on maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Language lists: ", languages)

	fmt.Println("JS: ", languages["JS"])

	delete(languages, "RB")

	fmt.Println("Language lists: ", languages)

	//loops are interesing in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("value is %v\n", value)
	}
}
