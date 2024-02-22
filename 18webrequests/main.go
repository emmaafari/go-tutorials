package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://gobyexample.com/"

func main() {
	fmt.Println("Welcome the the tutorial on web requests")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T", response)

	defer response.Body.Close() //Always close the connection

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println(content)
}
