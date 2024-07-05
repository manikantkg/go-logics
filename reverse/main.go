package main

import "fmt"

func main() {
	str := "Manikanta"
	fmt.Println(str)
	fmt.Println(Reverse(str))

}

func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
