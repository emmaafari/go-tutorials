package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://polded.com/products/NPML3KUWRQ?name=Emmauel&age=29"

func main() {
	fmt.Println("Welcome to the tutorial on URLs")

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams)
	for _, v := range qparams {
		fmt.Println("Param is: ", v)
	}

	partsIfUrl := &url.URL{
		Scheme:  "https",
		Host:    "polded.com",
		Path:    "/login",
		RawPath: "name=emma",
	}

	anotherURL := partsIfUrl.String()
	fmt.Println(anotherURL)
}
