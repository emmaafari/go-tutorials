package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emmaafari/mongoapi/route"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting starting...")

	r := route.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}
