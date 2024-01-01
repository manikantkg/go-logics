package main

import "fmt"

func main() {
	a, b := 0, 0

	// Initialize Value
	/*
		fmt.Printf("## INIT\n")
		fmt.Printf("Memory Location a: %p, b: %p\n", &a, &b)
		fmt.Printf("Value a: %d, b: %d\n", a, b) // 0 0
	*/
	// Passing By Value a(int)
	//Add(a) // Golang will copy value of 'a' and insert it into argument

	// Passing By Reference b(int), &b(*int) => with '&' we can get the memory location of 'b'
	AddPtr(&b) // Pass Memory Location of 'b' into argument

	fmt.Printf("\n## FINAL\n")
	fmt.Printf("Memory Location a: %p, b: %p\n", &a, &b)
	fmt.Printf("Value a: %d, b: %d\n", a, b) // 0 1
}

/*
// Pass By Value
func Add(x int) {
	fmt.Printf("\n## 'Add' Function\n")
	fmt.Printf("Before Add, Memory Location: %p, Value: %d\n", &x, x)
	x++
	fmt.Printf("After Add, Memory Location: %p, Value: %d\n", &x, x)
} */

// Pass By Reference
func AddPtr(x *int) {
	fmt.Printf("\n## 'AddPtr' Function\n")
	fmt.Printf("Before AddPtr, Memory Location: %p, Value: %d\n", x, *x)
	*x++ // We add * in front of the variable because it is a pointer, * will call value of a pointer
	fmt.Printf("After AddPtr, Memory Location: %p, Value: %d\n", x, *x)
}

//for different data types
/* func main() {
	// Slices
	fmt.Println("======================")
	fmt.Println("SLICES")
	fmt.Println("======================")
	var arrInt []int = []int{1, 2, 3, 4, 5}
	var sliceInt = arrInt[3:]

	fmt.Printf("Init\n")
	fmt.Printf("ArrInt: %+v, SliceInt: %+v\n\n", arrInt, sliceInt)

	sliceInt[0] = 10
	fmt.Printf("After\n")
	fmt.Printf("ArrInt: %+v, SliceInt: %+v\n", arrInt, sliceInt)

	// MAP
	fmt.Println("======================")
	fmt.Println("MAP")
	fmt.Println("======================")
	var emptyMap = make(map[string]interface{})
	fmt.Println("Init")
	fmt.Printf("emptyMap : %+v\n", emptyMap)

	MapFunc(emptyMap)
	fmt.Println("After")
	fmt.Printf("emptyMap : %+v\n", emptyMap)
}

func MapFunc(val map[string]interface{}) {
	val["this is a new value"] = 100
} */
