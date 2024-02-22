package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the tutorials on files")
	WriteFile()
	ReadFile("./myfile.txt")
}

func WriteFile() {
	content := "This needs to go in the file - learnCodeOnline.in "

	//Create the file
	file, err := os.Create("./myfile.txt")
	CheckNilError(err)

	//Write to the file
	length, err := io.WriteString(file, content)
	CheckNilError(err)

	fmt.Println("Length is: ", length)
	file.Close()
}

func ReadFile(filename string) {
	databyte, err := os.ReadFile(filename)
	CheckNilError(err)

	fmt.Println("Text Data inside the file is: \n", string(databyte))
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
