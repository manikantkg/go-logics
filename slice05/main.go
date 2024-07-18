package main

import (
	"fmt"
	"sort"
)

func main() {
	/* var things = []string{"abc", "bcd", "cde"}
	fmt.Println(things)
	fmt.Printf(" len %v and cap %v\n ", len(things), cap(things))

	things = append(things, "jkl")
	fmt.Println(things)
	fmt.Printf(" len %v and cap %v\n ", len(things), cap(things))

	things = append(things[1:])
	fmt.Println(things)
	fmt.Printf(" len %v and cap %v\n ", len(things), cap(things)) //Observe len and cap every time

	things = append(things[1 : len(things)-1])
	fmt.Println(things)

	heros := make([]string, 3, 3)
	heros[0] = "asdf"
	heros[1] = "qwer"
	heros[2] = "tyi"
	fmt.Printf("heros len %v and cap %v\n ", len(heros), cap(heros)) */

	/* heros = append(heros, "mnb")
	fmt.Printf(" heros len %v and cap %v\n ", len(heros), cap(heros))

	fmt.Println(heros)
	fmt.Println(cap(heros)) //observe the difference between cap & length

	myints := []int{3, 5, 7, 8}

	isSorted := sort.IntSlice(myints) //know more about sort library
	fmt.Println(isSorted, ">>>>>>>>>>")

	Arr() */
	arr1 := []int{1, 2, 4, 5, 6}
	arr2 := []int{9, 8, 4, 5, 7}
	res := Intersection(arr1, arr2)
	fmt.Println(">>>", res)

	//TestArray
	TestArray()
}

func Arr() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Arr is %v", a)
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:7]
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	m := []int{1, 2, 3, 4, 5}
	fmt.Println(m, &m)
	n := append(m[:2], m[3:]...)
	fmt.Println("##", n)
	fmt.Println("^^", m, &m)

}

func Intersection(arr1, arr2 []int) []int {
	intersection := make([]int, 0)
	for _, num1 := range arr1 {
		for _, num2 := range arr2 {
			if num1 == num2 {
				intersection = append(intersection, num1)
				break
			}
		}
	}
	return intersection
}

func TestArray() {
	var a []int = []int{1, 2, 3, 4, 5}
	fmt.Println("a is :", a)

	b := a
	fmt.Println("<><>", b)
	c := &a
	a[3] = 1000
	fmt.Println("c is : ", *c)

	var d []int
	d = append(d, a...)
	fmt.Println("d is ", d)

	//copy using loop
	newArray := make([]int, len(a))
	for i, v := range a {
		newArray[i] = v
	}
	fmt.Println("new array is: ", newArray)

	//using copy
	copiedArray := make([]int, len(d))
	copy(copiedArray, d)
	fmt.Println("copied Array : ", copiedArray)

	//copy slice elements by using append
	slice1 := []int{1, 2, 3}
	slice2 := []int{2, 5, 8}
	var s3 []int
	s3 = append(s3, slice1...)
	s3 = append(s3, slice2...)
	fmt.Println("resilt is >>> ", s3)
	sort.Ints(s3)
	fmt.Println(s3)
}
