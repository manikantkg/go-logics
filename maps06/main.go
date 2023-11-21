package main

import "fmt"

func main() {

	// make && new

	// new  -> only allocates - no init of memory
	// make -> allocation and init - non zerod storage

	// var score map[string]int
	// score["mani"] = 89

	// fmt.Println(score)     //panic

	score := make(map[string]int)
	score["mani"] = 89
	score["karthikeya"] = 67
	score["manu"] = 87
	score["chinnodu"] = 56

	getScore := score["chinnodu"]
	fmt.Println(getScore, "<><><><><<>")

	delete(score, "chinnodu")

	fmt.Println(score)

	for k, v := range score {
		fmt.Printf("Score of %v and value is %v\n", k, v)
	}

}
