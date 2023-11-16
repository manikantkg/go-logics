package main

import "fmt"

//Count number of capital numbersx

type deck []string

func (c deck) CapitalNumbers(s string) {
	var counts int = 0
	for _, v := range s {
		if v >= 65 && v <= 90 {
			counts++
		}
	}
	fmt.Println(counts)
}

// func main() {
// 	CapitalNumbers("My Name Is Manikanta")

// }
