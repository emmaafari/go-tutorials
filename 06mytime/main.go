package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study in golang")

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("2006-01-02 15:04:05 Monday"))

	createdDate := time.Date(presentTime.Year(), time.February, 8, 16, 0, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
