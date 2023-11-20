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
	// 	fmt.Printf("index is %v and value is %v\n", v)
	// }

	rawValue := 1

	for rawValue < 10 {

		if rawValue == 5 {
			rawValue++
			break //continue also works but understand the difference between break and continue
		}
		fmt.Println("Value is ", rawValue)
		rawValue++
	}

}
