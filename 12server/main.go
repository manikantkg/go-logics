package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welocme to web API requests")

	// PerformGetRequest()
	// PerforPostJsonRequest()
	PerformPostFormGetRequest()
}

func PerformGetRequest() {

}

func PerforPostJsonRequest() {

}

func PerformPostFormGetRequest() {
	const myurl = "http://localhost:8000/get"

	//formdata
	data := url.Values{}
	data.Add("firstname", "mani")
	data.Add("lastname", "kanta")
	data.Add("email", "mani@123.com")

	resp, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	// resp, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }

	defer resp.Body.Close()

	fmt.Println("Status code : ", resp.StatusCode)
	fmt.Println("Content length: ", resp.ContentLength)

	// var respString strings.Builder

	content, _ := ioutil.ReadAll(resp.Body)
	// byteCount, _ := respString.Write(content)
	fmt.Println(string(content))
	// fmt.Println("ByteCount", byteCount)
}
