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
	//printing indidvidual elements
	fmt.Println("First Person Details:")
	fmt.Println("Name:", p[0].Name)
	fmt.Println("Location:", p[0].Loc)
	fmt.Println("State:", p[0].State)
	fmt.Println() // Print a blank line for separation

	// Print the last person's details
	fmt.Println("Last Person Details:")
	fmt.Println("Name:", p[len(p)-1].Name)
	fmt.Println("Location:", p[len(p)-1].Loc)
	fmt.Println("State:", p[len(p)-1].State)
	//
	mani := User{"Manikanta", 33, "manikanta@123.com", true}
	//fmt.Println(mani)
	fmt.Printf("Mani detials are : %+v\n", mani)
	fmt.Printf("Mani name is  %v and Email is: %v", mani.Name, mani.Email)
}
