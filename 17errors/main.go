package main

import (
	"errors"
	"fmt"
)

func main() {
	message := "Hello!!!"
	myError := errors.New("Wrong message") //error.New() function added here
	if message != "Programe" {
		fmt.Println(myError)
	}

	FormatEroor()
	Custom()
}

//2)Errorf using fmt library

func FormatEroor() {
	var b, a int
	b = 0
	a = 4
	error := fmt.Errorf("%d is lesser than a", b)
	if a > 0 && b != 0 {
		fmt.Println("a is  divisible", b)
	} else {
		fmt.Println(error)
	}

}

//custom error

type DivisionByZero struct {
	message string
}

// define Error() method on the struct
func (z DivisionByZero) Error() string {
	return "Number Cannot Be Divided by Zero"
}

func divide(n1 int, n2 int) (int, error) {

	if n2 == 0 {
		return 0, &DivisionByZero{}
	} else {
		return n1 / n2, nil
	}
}

func Custom() {

	number1 := 15
	// change the value of number2 to get different result
	number2 := 0

	result, err := divide(number1, number2)

	// check if error occur or not
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %d", result)
	}
}
