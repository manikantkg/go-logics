package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func readData(ch <-chan int) {
	x := <-ch // Receiving data from the channel
	fmt.Println(x)
}
func main() {

	ch := make(<-chan int) // Declaring a receive-only channel

	go readData(ch)

	/*
		fmt.Println("Starts from here!!!!")
		wg.Add(2)
		p := make(chan int)
		//p <- 10

		
		go Addition(p)
		go Sub(p)
		wg.Wait()
		final := <-p
		fmt.Println("ended", final) */

	/* ch := make(chan int) // Creating an unbuffered channel

	go func() {
		time.Sleep(time.Second) // Simulating some work
		ch <- 5                 // Sending a value to the channel
	}()

	x := <-ch      // Receiving the value from the channel
	fmt.Println(x) // Output: 5 */
}

func Addition(p chan int) {

	defer wg.Done()

	p <- 10
	var a, b int = 5, 5

	c := a + b
	fmt.Println(c, p)
}

func Sub(p chan int) {

	defer wg.Done()

	p <- 10
	var a, b int = 10, 5
	c := a - b
	fmt.Println(c, p)

}
