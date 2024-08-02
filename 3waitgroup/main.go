package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	defer wg.Done()
	//
	wg.Add(2)
	go hello()
	go world()
	wg.Done()

	fmt.Println("Mani ends!!!")
	//using anonymus
	wg.Add(11)
	abc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range abc {
		go func(i int) {
			wg.Wait()
			fmt.Println("i is :", i, "index value is >>", v)
		}(v)
		// fmt.Println(i, v)
	}
}

func hello() {
	defer wg.Done()
	fmt.Println("Hello")

}
func world() {
	defer wg.Done()
	fmt.Println("World")
}
