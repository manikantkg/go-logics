package main

import "fmt"

func main() {
	fmt.Println("Execution of defer!!") //LIFO -> Last in First OUT

	defer fmt.Println("One >>> 1")
	defer fmt.Println("Two >>> 2")
	defer fmt.Println("Three >> 3")
	Abc()

}

func Abc() {
	fmt.Println("Its a defer function")
	a := []int{12, 23, 34, 45, 57}
	fmt.Println(a)

	for _, ab := range a {
		fmt.Println(ab)
	}
}
