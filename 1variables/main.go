package main

import "fmt"

func main() {
	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string
	superman = "I am superman"
	fmt.Println(superman)

	thor := "I am thor"
	fmt.Println(thor)

	thorRating := 3
	fmt.Printf("thorRating is %v and type %T", thorRating, thorRating)

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
	fmt.Println(spiderman, age, powers, antman)
}
