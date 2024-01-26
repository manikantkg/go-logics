package main

import (
	"fmt"
	"time"
)

/*
	 func Add(x, y int) {
		C := x + y
		fmt.Println("c is :", C)
	}
*/
func sum(p chan int) {
	var b int = 23
	//p = make(chan int)
	a := <-p
	c := a + b
	fmt.Println("c is:", c)
}

func main() {
	/* fmt.Println("Goroutineds Stats from here!!!")
	go Add(3, 5)                //simple go routine
	time.Sleep(2 * time.Second) //holds the main function execution by using time package
	fmt.Println("Goroutineds Ends from here!!!") */

	fmt.Println(">>>>>>>>>>>>>>")

	d := make(chan int)
	//fmt.Println(d)
	d <- 3
	panic("Error occured")
	fmt.Println(<-d)
	go sum(d)
	time.Sleep(2 * time.Second)

}
