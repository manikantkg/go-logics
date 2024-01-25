package main

import "fmt"

func main() {
	// defer fmt.Println("one") //observe the output
	// fmt.Println("Two")
	// fmt.Println("Three")
	// Deferfunction()
	PanicFunction(2, 4)
	PanicFunction(2, 0)

}

func Deferfunction() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

//panic usage

func hadlepanic() {
	a := recover()
	if a != nil {
		fmt.Println("Recover", a)
	}
}

func PanicFunction(a, b int) {
	defer hadlepanic()
	div := a / b
	fmt.Println("div is :", div)
	panic("b value is lesser than A")

	// fmt.Println("Help! Something bad is happening.")
	// panic("Ending the program")
	// fmt.Println("Waiting to execute")
}
