package main

import (
	"fmt"
	"time"
)

// func main() {
// 	fmt.Println("Channels in Golang!!!")

// 	msg := make(chan string)

// 	go greet(msg)
// 	time.Sleep(2 * time.Second)

// 	greeting := <-msg
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Greeting received")
// 	fmt.Println(greeting)

// }

func main() {
	msg := make(chan string)

	go greet(msg)
	time.Sleep(2 * time.Second)

	greeting := <-msg
	time.Sleep(2 * time.Second)
	fmt.Println("Greeting received")
	fmt.Println(greeting)

	_, ok := <-msg
	if ok {
		fmt.Println("Channel is open")
	} else {
		fmt.Println("Channel is closed")
	}
}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!!")

	ch <- "Hai Manikanta"

	close(ch) // Use this for 2nd main function

	fmt.Println("Greeter Completed")

}
