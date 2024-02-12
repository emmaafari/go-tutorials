package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome the user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T ", input)
}
