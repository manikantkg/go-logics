package main

import "fmt"

func main() {
	defer fmt.Println("one") //observe the output
	fmt.Println("Two")
	fmt.Println("Three")
	Deferfunction()
}

func Deferfunction() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}
