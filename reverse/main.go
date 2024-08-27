package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Manikanta"
	fmt.Println(str)
	fmt.Println(Reverse(str))
	fmt.Println(Rev(str))

	//Fibino
	printFibonacciSeries()

	//Prime
	PrimeNumber()

	//string reverse infosys
	StringReverse("mani kan ta")

	//single string reverse
	singleStringReverse()

	//printing excel columns

	var numTests int
	fmt.Println("Enter the number of test cases:")
	fmt.Scan(&numTests)

	tests := make([]int, numTests)
	fmt.Println("Enter the numbers:")
	for i := 0; i < numTests; i++ {
		fmt.Scan(&tests[i])
	}

	for _, n := range tests {
		fmt.Printf("n = %d ---> %s\n", n, NumberToExcelColumn(n))
	}

}

// revwersing the characters in the same place
func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

// Reversing the srings
func Rev(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// Reversing the string using string library
func StringReverse(s string) {
	words := strings.Fields(s)
	for i, word := range words {
		r := []rune(word)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		words[i] = string(r)
	}
	fmt.Println(strings.Join(words, " "))

}

// single string reverse
func singleStringReverse() {
	var num string = "abcde"
	r := []rune(num)

	for i := 0; i < len(r)/2; i++ {
		temp := r[i]
		r[i] = r[len(r)-1-i]
		r[len(r)-1-i] = temp
	}

	fmt.Println(string(r))
}

//printFibonacciSeries

func printFibonacciSeries() {
	n := 15
	a, b := 0, 1
	c := b
	for true {
		c = b
		b = a + b
		if b >= n {
			fmt.Println()
			break
		}
		a = c
		fmt.Println(b)

	}
}

//Prime NUmber

func PrimeNumber() {
	var num, count int
	count = 0
	fmt.Println("Enter the Number")
	fmt.Scanln(&num)
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			count++
			break
		}
	}
	if count == 0 && num != 1 {
		fmt.Println("num is prime", num)
	} else {
		fmt.Println("num is not prime", num)
	}
}

// Printing column numbers
func NumberToExcelColumn(n int) string {
	column := ""
	for n > 0 {
		n-- // Adjust because we want 1-based index for letters
		column = string(rune('A'+(n%26))) + column
		n /= 26
	}
	return column
}
