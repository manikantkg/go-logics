package main

import "fmt"

func main() {
	Fact(4)

}

func Fact(n int) {
	var a int = 1
	for i := 1; i <= n; i++ {
		a *= i
		fmt.Println(a)
	}

}
