package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func main() {
	mani := User{"Manikanta", 33, "manikanta@123.com", true}
	//fmt.Println(mani)
	fmt.Printf("Mani detials are : %+v\n", mani)
	fmt.Printf("Mani name is  %v and Email is: %v", mani.Name, mani.Email)

}
