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

	//Mutiple function
	_, b, _ := Multiple(23, "mani", 5.9)
	fmt.Println("Fetch value: ", b)
	//second method
	res := Multiple2(34, "manikanta", 5.6)
	fmt.Println("REad value: ", res.StringValue)
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

//Returning multiple things with escape operator
func Multiple(a int, b string, c float32) (int, string, float32) {
	return a, b, c
}

//second method for fetch only one output from multiple return values

type Values struct {
	IntValue    int
	StringValue string
	FloatValue  float32
}

func Multiple2(a int, b string, c float32) Values {
	return Values{
		IntValue:    a,
		StringValue: b,
		FloatValue:  c,
	}
}
