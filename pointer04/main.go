package main

import "fmt"

func main() {
	var score int = 34

	// var p *int

	// if p != nil {
	// 	fmt.Println("P is having value is : ", *p)
	// } else {
	// 	fmt.Println("Watchout for nil value")
	// }
	p := &score
	fmt.Println(p)

	// var lifeBooster float64 = 99.2
	// lifeBoosterRef := &lifeBooster

	// fmt.Println(lifeBooster)
	// fmt.Println(lifeBoosterRef)
	// fmt.Println(*lifeBoosterRef)

	var numbers [3]string
	numbers[0] = "abc"
	numbers[1] = "sdf"

	numbers[2] = "abwerc"
	fmt.Println(numbers)

	var colors = [4]string{"red", "grey", "ash", "blue"}
	fmt.Println(colors)
	fmt.Println(colors[3])
	fmt.Println(len(colors))

}
