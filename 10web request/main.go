package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.mpstdc.com/" //"https://lco.dev" //"https://google.co.in"

func main() {
	fmt.Println("WeB request")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Resp is of type: %T\n", resp)
	// fmt.Println("respoinse is >>", resp)

	defer resp.Body.Close() //caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
