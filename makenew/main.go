package main

import "fmt"

func main() {

	a := make([]int, 3)
	fmt.Println(a, len(a), cap(a), &a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(&b[1])

}
