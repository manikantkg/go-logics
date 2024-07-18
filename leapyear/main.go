package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number ")
	fmt.Scanln(&n)
	if n%4 == 0 && n/100 != 0 || n%400 == 0 {
		fmt.Println("Leap year")
	} else {
		fmt.Println("Not Leap year")
	}
}
