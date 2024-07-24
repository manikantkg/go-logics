package main

import (
	"fmt"
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

func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func Rev(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
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
	fmt.Println("Rnter the Number")
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
