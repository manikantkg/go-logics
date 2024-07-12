package main

import "fmt"

type User struct {
	Name   string //must start with Capital name only otherwise it wont export
	Age    int
	Email  string
	Status bool
}

//No inheritence && treated as class or parent
type Person struct {
	Name  string
	Loc   string
	State string
}

func main() {

	// var p []Person
	p := []Person{
		{
			Name:  "ABC",
			Loc:   "HYD",
			State: "TS",
		},
		{Name: "BC",
			Loc:   "HYD",
			State: "TS",
		},
		{
			Name:  "ABC",
			Loc:   "HYD",
			State: "TS",
		},
	}

	for _, person := range p {
		fmt.Println(person)
	}

	mani := User{"Manikanta", 33, "manikanta@123.com", true}
	//fmt.Println(mani)
	fmt.Printf("Mani detials are : %+v\n", mani)
	fmt.Printf("Mani name is  %v and Email is: %v", mani.Name, mani.Email)
}
