package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the tutorial on making a request")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFromRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	// fmt.Println(string(content))
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count:", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Golang",
			"price":100,
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFromRequest() {
	const myUrl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "Emmanuel")
	data.Add("lastname", "Kofi")
	data.Add("email", "emma@code.dev")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
