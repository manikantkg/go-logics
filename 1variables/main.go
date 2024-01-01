package main

import (
	"fmt"
	"reflect"
)

func main() {
	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string
	superman = "I am superman"
	fmt.Println(superman)

	thor := "I am thor"
	fmt.Println(thor)
	fmt.Println(reflect.TypeOf(thor)) //finding type with the help of reflect package
	// fmt.Println(reflect.Interface, "thor")

	thorRating := 3
	fmt.Printf("thorRating is %v and type %T \n", thorRating, thorRating)

	var Itonman, CapAmerica string = "I am Ironman", "I am captain America"
	fmt.Println(Itonman, CapAmerica)

	var defaultValue string
	fmt.Println(defaultValue)

	var (
		spiderman = "I am a Spiderman"
		age       = 33
		powers    = "web splings, spidy sense"
		antman    = "I am antman"
	)
	fmt.Println(spiderman, "age is :", age, powers, antman)

	var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)
}
