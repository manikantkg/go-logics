package main

import (
	"fmt"
	"sort"
)

func main() {

	// Prinitng Unique elements
	// arr1 := []int{1, 3, 4, 5, 7, 4, 2, 2, 4, 5}
	// fmt.Println("arr1>>", arr1)
	// Unique_items := Unique(arr1)
	// fmt.Println(Unique_items)
	//---------------------------------

	//Removing Susequences

	// abc := Remove("aaaabnddkemkdflenkd")
	// fmt.Println(abc)

	//Rem
	// Rem()

	//Capital
	// CapitalNumbers("My Name Is Manikanta")

	//Maps
	KeysPrint()
	Values()
}

// Prinitng Unique elements
func Unique(arr []int) []int {
	repeated := map[int]bool{}
	result := []int{}
	for i := range arr {
		if repeated[arr[i]] != true {
			repeated[arr[i]] = true
			result = append(result, arr[i])
		}
	}
	return result
}

//Removing Susequences

func Remove(s string) string {
	c := []rune(s)
	for i := len(c) - 1; i > 0; i-- {
		if i < len(c) && c[i] == c[i-1] {
			c = append(c[:i-1], c[i+1:]...)
		}
	}
	return string(c)
}

//Removing duplicates

func Rem() {
	arry := []int{1, 2, 3, 4, 5, 5, 6, 7, 4, 3}
	fmt.Println(arry)

	for i := 0; i < len(arry); i++ {
		for j := 0; j < len(arry); j++ {
			if i == j {
				continue
			}
			if arry[i] == arry[j] {
				arry = append(arry[:j], arry[j+1:]...)
			}
		}
	}
	fmt.Println(arry)
}

//Count number of capital numbers

func CapitalNumbers(s string) {
	var count int = 0
	for _, v := range s {
		if v >= 65 && v <= 90 {
			count++
		}
	}
	fmt.Println(count)
}

//sort maps by value

func Values() {
	basket := map[string]int{"orange": 5, "apple": 7,
		"mango": 3, "strawberry": 9}

	keys := make([]string, 0, len(basket))

	for key := range basket {
		keys = append(keys, key)
	}

	fmt.Println(basket)
	fmt.Println(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return basket[keys[i]] < basket[keys[j]]
	})

	fmt.Println(keys)
}

//Sort map by keys

func KeysPrint() {
	basket := map[string]int{"orange": 5, "apple": 7,
		"mango": 3, "strawberry": 9}

	keys := make([]string, 0, len(basket))

	for k := range basket {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, basket[k])
	}
}
