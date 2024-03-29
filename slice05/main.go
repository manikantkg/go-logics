package main

import (
	"fmt"
	"sort"
)

func main() {
	var things = []string{"abc", "bcd", "cde"}
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
	fmt.Printf("heros len %v and cap %v\n ", len(heros), cap(heros))

	heros = append(heros, "mnb")
	fmt.Printf(" heros len %v and cap %v\n ", len(heros), cap(heros))

	fmt.Println(heros)
	fmt.Println(cap(heros)) //observe the difference between cap & length

	myints := []int{3, 5, 7, 8}

	isSorted := sort.IntSlice(myints) //know more about sort library
	fmt.Println(isSorted, ">>>>>>>>>>")

	Arr()
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
