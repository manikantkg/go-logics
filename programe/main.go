package main

import "fmt"

func main() {
	//Fact(4)
	//numbers()
	//pyramid()
	ReversePyramid()

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

//Reverse pyramid
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
