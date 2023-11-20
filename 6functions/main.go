package main

import "fmt"

func main() {
	fmt.Println("Functions started in Golang")
	greeter()

	result := adder(3, 5)
	fmt.Println(result)
	proRes, myMessage := proAdder(2, 5, 7, 8)
	fmt.Println(proRes, "proResult is ")
	fmt.Println(myMessage)

}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, v := range values {
		total += v
	}

	return total, "HI Pro result function"
}

func greeter() {
	fmt.Println("Good Evening Golang ^^^")

}
