package main

import (
	"fmt"
	"sort"
)

func main() {

	//Removing Susequences

	// abc := Remove("aaaabnddkemkdflenkd")
	// fmt.Println(abc)

	//Rem
	// Rem()

	//Capital
	//deck.CapitalNumbers(deck{}, "My Name Is Manikanta")

	//Maps
	// KeysPrint()
	// Values()

	//TraingleNumbers()

	// Pyramid()

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
	fmt.Println(result)
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

//sort maps by values

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
	basket := map[string]int{
		"orange":     5,
		"apple":      7,
		"mango":      3,
		"strawberry": 9}

	keys := make([]string, 0, len(basket))

	for k := range basket {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, basket[k])
	}
}

//print Traingular Numbers

func TraingleNumbers() {
	var r int = 5
	for i := 1; i <= r; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}

}

// Pyramid starts
func Pyramid() {

	var rows int = 5
	var k int
	for i := 1; i <= rows; i++ {
		k = 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}
}

/*
	// r, temp := 5, 1
	// 	for i := 0; i < 5; i++ {
	// 		for j := 1; j < r-1; j++ {
	// 			fmt.Print("")
	// 		}
	// 		for k := 0; k <= i; k++ {
	// 			if k == 0 || i == 0 {
	// 				temp = 1
	// 			} else {
	// 				temp = temp * (i - k + 1) / k
	// 			}
	// 			fmt.Printf("%d ", temp)
	// 		}
	// 		fmt.Println(" ")
	// 	}
	// 	fmt.Println(r, temp)
}
*/
