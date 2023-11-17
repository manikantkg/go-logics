package main

import "fmt"

// if-else Condition implemented
// map initialization implementation
func main() {
	/* var number int = 10
	var result string

	if number < 10 {
		result = "EMI's need to pay"
	} else if number > 10 {
		result = "Need to check pending EMI's"
	} else {
		result = "End date to pay EMIs"
	}
	fmt.Println(result) */

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

	for i, _ := range a {
		fmt.Println(a[i])
	}
	fmt.Println(">>>>>>>>>>>", a[3])

	/* var num int
	fmt.Println("Enter the Number")
	fmt.Scanln(&num)
	if num >= 1 && num <= 5 {
		fmt.Println("Mid of the week")
	} else if num >= 6 {
		fmt.Println("Weekend!!")
	} */

}
