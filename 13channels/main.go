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
	//Select function
	Select()
}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!!")

	ch <- "Hai Manikanta"

	close(ch) // Use this for 2nd main function

	fmt.Println("Greeter Completed")

}

func Select() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	default:
		fmt.Println("No messages received")
	}
}
