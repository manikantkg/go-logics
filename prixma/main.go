package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup

func printEven(evenChan, oddChan chan int) {
	defer wg.Done()
	for i := 0; i < 10; i += 2 {
		fmt.Println("Even:", i)
		evenChan <- i
		<-oddChan
	}
}

func printOdd(evenChan, oddChan chan int) {
	defer wg.Done()
	for i := 1; i < 10; i += 2 {
		fmt.Println("Odd:", i)
		oddChan <- i
		<-evenChan
	}
}

func main() {
	/* evenChan := make(chan int)
	oddChan := make(chan int)

	wg.Add(2)

	go printEven(evenChan, oddChan)
	go printOdd(evenChan, oddChan)

	wg.Wait()
	close(evenChan)
	close(oddChan) */
	Even()
	Odd()
}

func Even() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even ",i)
		}

	}
}

func Odd() {
	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("Odd ", i)
		}

	}
}
