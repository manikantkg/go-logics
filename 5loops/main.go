package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	// days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	// fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for _, v := range days {
	// 	fmt.Printf("index is  and value is %v\n", v)
	// }

	rowValue := 1

	for rowValue < 10 {

		if rowValue == 5 {
			goto lco
		}

		if rowValue == 5 {
			rowValue++
			break //continue also works but understand the difference between break and continue
		}
		fmt.Println("Value is ", rowValue)
		rowValue++
	}

lco:
	fmt.Println("Jumping learn Golang")
}
