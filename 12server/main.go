package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welocme to web API requests")

	PerformGetRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:8000/get"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status code : ", resp.StatusCode)
	fmt.Println("Content length: ", resp.ContentLength)

	var respString strings.Builder

	content, _ := ioutil.ReadAll(resp.Body)
	byteCount, _ := respString.Write(content)
	// fmt.Println(string(content))
	fmt.Println("ByteCount", byteCount)
}
