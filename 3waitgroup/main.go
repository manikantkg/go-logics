package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(11)
	abc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, _ := range abc {

		go func(i int) {
			defer wg.Done()
			fmt.Println("i is :", i, "index value is >>", abc[i])
		}(i)

		// fmt.Println(i, v)
	}
	wg.Wait()
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go hello(&wg)
// 	go world()
// 	wg.Done()

// 	fmt.Println("Mani ends!!!")

// }

// func hello(wg sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Hello")

// }
// func world(wg sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("World")
// }
