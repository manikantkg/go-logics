package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	abc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, _ := range abc {

		go func(i int) {
			defer wg.Done()
			fmt.Println(i, abc[i])
		}(i)

		// fmt.Println(i, v)
	}
	wg.Wait()
}
