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

	//is Unique
	fmt.Println(isUnique("abcdef")) // true
	fmt.Println(isUnique("aabbcc")) // false
	//FizzBuzz
	num := 100
	FizzBuzz(num)
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

//is unique RXbenifits question2
func isUnique(s string) bool {
	// Assuming ASCII characters, there are only 128 unique possible characters.
	if len(s) > 128 {
		return false
	}

	// Map to keep track of characters we have seen.
	charMap := make(map[rune]bool)

	for _, char := range s {
		if charMap[char] {
			return false // Character already seen, hence not unique.
		}
		charMap[char] = true
	}
	return true
}

//FizzBuzz RXbenifits Question1

func FizzBuzz(num int) {
	for i := 1; i < num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
