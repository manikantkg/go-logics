package main

import "fmt"

type User struct {
	Name   string //must start with Capital name only otherwise it wont export
	Age    int
	Email  string
	Status bool
}

//No inheritence && treated as class or parent
func main() {
	mani := User{"Manikanta", 33, "manikanta@123.com", true}
	//fmt.Println(mani)
	// fmt.Printf("Mani detials are : %+v\n", mani)
	// fmt.Printf("Mani name is  %v and Email is: %v", mani.Name, mani.Email)
	mani.GetStatus()
	mani.NewMail()

}

func (u User) GetStatus() {
	fmt.Println("Is user Active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@123.com"
	fmt.Println("EMail of User is : ", u.Email)
}
