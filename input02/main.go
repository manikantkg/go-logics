package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// var myString string
	// fmt.Scanln(&myString)
	// fmt.Println(myString)

	// var name string = "Manikanta"
	// var a, _ = fmt.Println(name)
	// fmt.Println(a)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your full name: ")
	// myName, _ := reader.ReadString('\n')
	// fmt.Println(myName)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your rating: ")
	myRating, _ := reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64) //converting
	fmt.Println(mynumRating + 2)
}
