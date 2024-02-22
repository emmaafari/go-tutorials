package main

import "fmt"

func main() {
	fmt.Println("Welcome to the tutorial on methods")

	user := User{
		"Emmanuel Kofi",
		"emmauel.kofi@gmail.com",
		true,
		29,
	}

	fmt.Println("user: ", user)
	fmt.Printf("user details are: %+v\n", user)
	fmt.Printf("Name is %v and email is %v\n", user.Name, user.Email)

	user.GetStatus()
	user.NewMail()

	fmt.Printf("Name is %v and email is %v\n", user.Name, user.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active?", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
