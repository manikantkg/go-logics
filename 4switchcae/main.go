package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch Case Implementation")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("dice number value is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot ")
	case 3:
		fmt.Println("you can move 3 spot ")
	case 4:
		fmt.Println("you can move 4 spot ")
	case 5:
		fmt.Println("you can move 2 spot ")
	case 6:
		fmt.Println("You can move 6 spot && roll again")
	default:
		fmt.Println("What is what!!")
	}
}
