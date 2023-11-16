package main

import "fmt"

func UniqueElement() {
	// Prinitng Unique elements
	arr1 := []int{1, 3, 4, 5, 7, 4, 2, 2, 4, 5}
	fmt.Println("arr1>>", arr1)
	Unique_items := Unique(arr1)
	fmt.Println(Unique_items)

}
