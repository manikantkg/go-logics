package main

import (
	"fmt"
	"sort"
)

func main() {
	var things = []string{"abc", "bcd", "cde"}
	fmt.Println(things)

	things = append(things, "jkl")
	fmt.Println(things)

	things = append(things[1:])
	fmt.Println(things)

	things = append(things[1 : len(things)-1])
	fmt.Println(things)

	heros := make([]string, 3, 3)
	heros[0] = "asdf"
	heros[1] = "qwer"
	heros[2] = "tyi"

	heros = append(heros, "mnb")
	fmt.Println(heros)
	fmt.Println(cap(heros)) //observe the difference between cap & length

	myints := []int{3, 5, 7, 8}

	isSorted := sort.IntSlice(myints) //know more about sort library
	fmt.Println(isSorted, ">>>>>>>>>>")
}
