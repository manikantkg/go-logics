package main

import (
	"fmt"
	"time"
)

func Add(x, y int) {
	C := x + y
	fmt.Println("c is :", C)
}

func Sum(a chan int) {
	var b, c int
	a <- a
	c = a + b
	fmt.Println("c is:", c)
}

func main() {
	fmt.Println("Goroutineds Stats from here!!!")
	go Add(3, 5)                //simple go routine
	time.Sleep(2 * time.Second) //holds the main function execution by using time package
	fmt.Println("Goroutineds Ends from here!!!")

}
