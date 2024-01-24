package main

import (
	"fmt"
	"sort"
)

func main() {
	//Fact(4)
	//numbers()
	//pyramid()
	//ReversePyramid()
	//SortSlice()
	// sortArray()
	Amstrong()

}

func Fact(n int) {
	var a int = 1
	for i := 1; i <= n; i++ {
		// a *= i
		a = a * i
		fmt.Println(a)
	}

}

func numbers() { //need to check the output
	for i := 1; i <= 10; i++ {
		i *= i
		fmt.Println(i)

	}
}

func pyramid() {
	rows := 5

	for i := 1; i < rows; i++ {
		for j := 1; j < rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k != (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

// Reverse pyramid
func ReversePyramid() {
	var rows int = 5

	for i := rows; i >= 1; i-- {
		for space := 1; space <= rows-i; space++ {
			fmt.Print(" ")
		}
		for j := i; j <= 2*i-1; j++ {
			fmt.Printf("*")
		}
		for j := 0; j < i-1; j++ { //j:=i half pyramid will come
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}

//sorting slice

func SortSlice() {
	a := []int{5, 3, 4, 7, 8, 9}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j] // a[i] > a[j] for descending order
	})
	for _, v := range a {
		fmt.Println(v)
	}

}

// sorting with loops only
func sortArray() {
	arr := [5]int{50, 30, 20, 10, 40}
	var min, temp int
	for i := 0; i <= 4; i++ {
		min = i
		for j := i + 1; j <= 4; j++ {
			if arr[j] < arr[min] {

				// changing the index to show the min value
				min = j
			}
		}
		temp = arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
	fmt.Println(arr)
}

// Amstrong number
func Amstrong() {
	fmt.Println("Number = 153")
	// declare the variables
	var number, temp, remainder int
	var result int = 0
	// initialize the variables
	number = 153
	temp = number
	// Use of For Loop
	for {
		remainder = temp % 10
		result += remainder * remainder * remainder
		temp /= 10
		if temp == 0 {
			break // Break Statement used to stop the loop
		}
	}
	// If satisfies Armstrong condition
	if result == number {
		fmt.Printf("%d is an Armstrong number.", number)
	} else {
		fmt.Printf("%d is not an Armstrong number.", number)
	}
	// print the result
}
