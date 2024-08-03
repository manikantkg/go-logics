//str:="mainkanta"
//substring:= nta

package main

import "fmt"

func main() {

	str := "manikanta"
	substr := "nta"

	fmt.Println(Contain(str, substr))

}
func Contain(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		temp := str[i : i+len(substr)]
		fmt.Println(temp)
		if temp == substr {
			return true
		}

	}
	return false
}
