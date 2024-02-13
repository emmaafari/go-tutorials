package main

import "fmt"

func main() {
	fmt.Println("Welcome to the tutorial on structs")

	user := User{
		"Emmanuel Kofi",
		"emmauel.kofi@gmail.com",
		true,
		29,
	}

	fmt.Println("user: ", user)
	fmt.Printf("user details are: %+v\n", user)
	fmt.Printf("Name is %v and email is %v", user.Name, user.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
