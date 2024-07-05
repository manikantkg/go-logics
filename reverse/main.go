package main

import "fmt"

func main() {
	str := "Manikanta"
	fmt.Println(str)
	fmt.Println(Reverse(str))
	fmt.Println(Rev(str))

}

func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func Rev(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
