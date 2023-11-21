package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	content := "This needs to go in a file"

	file, err := os.Create("./go-locgics.txt")
	checkNilErr(err) //instead of writing below code you can write a commom function
	// if err != nil {
	// 	panic(err)
	// }
	length, err := io.WriteString(file, content)
	checkNilErr(err) //instead of writing below code you can write a commom function
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./go-locgics.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	checkNilErr(err)
	//instead of writing below code you can write a commom function
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Text inside the data is \n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
