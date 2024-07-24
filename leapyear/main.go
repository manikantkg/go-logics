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
	//count upper case function
	CountUpperCase()

	//Factors
	Factors()
}

//count uppercase letters
func CountUpperCase() {

	a := "My Name Is ManiMantA" //M-N-I-M-M-A
	var count int = 0
	for _, v := range a {
		if v >= 65 && v <= 90 {
			count++
		}
	}
	fmt.Println("Count is : ", count)
}

//Factors of a number
func Factors() {
	var num int
	fmt.Print("Enter the num: ")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}

	}
}
