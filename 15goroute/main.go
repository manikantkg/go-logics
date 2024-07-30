package main

import (
	"fmt"
	"sync"
)

/*
	 func Add(x, y int) {
		C := x + y
		fmt.Println("c is :", C)
	}
*/
/* func sum(p chan<- int) {
	// var p chan int
	var b int = 23

	p <- 3
	c := p + b
	fmt.Println("c is:", c, p)
} */

func main() {
	/* fmt.Println("Goroutineds Stats from here!!!")
	go Add(3, 5)                //simple go routine
	time.Sleep(2 * time.Second) //holds the main function execution by using time package
	fmt.Println("Goroutineds Ends from here!!!") */

	/* fmt.Println(">>>>>>>>>>>>>>")

	p := make(chan int)
	//fmt.Println(d)
	// d <- 3
	// panic("Error occured")
	//fmt.Println(<-d)
	go sum(p)
	time.Sleep(2 * time.Second) */

	ch := make(chan int)
	go sum(ch)
	result := <-ch
	fmt.Println(">>>>", result)

	//waitgroups
	c := make(chan int)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go R(c)

	}
	go func() {  //if we dont use this anonymus function we will get deadlock condition
		wg.Wait()
		close(c)
	}()

	for range c {
		// This loop will continue until the channel is closed
	}

	fmt.Println("All goroutines finished executing")

}

func sum(ch chan<- int) {
	sum, n := 0, 10
	for i := 1; i <= n; i++ {
		sum += i
		fmt.Println("sum is ", sum)
	}
	ch <- sum // Send the sum to the channel
}

//wait groups

var wg sync.WaitGroup

func R(c chan int) {
	defer wg.Done()
	a := 12
	b := 10
	d := b + a
	fmt.Println(d)
	c <- 1

}
