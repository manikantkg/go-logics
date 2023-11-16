package main

import "fmt"

//if-else Condition implemented
//map initialization implementation
func main() {
	var number int = 10
	var result string

	if number < 10 {
		result = "EMI's need to pay"
	} else if number > 10 {
		result = "Need to check pending EMI's"
	} else {
		result = "End date to pay EMIs"
	}
	fmt.Println(result)

	var num int
	var a = map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
	fmt.Println(a)
	fmt.Println("Enter the Number")
	fmt.Scanf("Enter the Number is : ", &num)

	if num == 1 || num == 2 || num == 3 || num == 4 || num == 5 {
		fmt.Println("Mid of the week")
	} else {
		fmt.Println("Weekend!!")
	}

}
