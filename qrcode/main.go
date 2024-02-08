package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/skip2/go-qrcode"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int ``
	Name     string
	Age      int
	Location string
}

func main() {

	//qrcode
	dsn := "root:ap02BL1426*@tcp(127.0.0.1:3306)/qrcode?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("ERROR while connecting DB>>")
	} else {
		fmt.Println("Connected successfully!!!")
	}

	db.AutoMigrate(&User{})

	//user := User{Name: "Chinnodu", Age: 18, Location: "Hyderabad"}

	//db.Create(&user) // pass pointer of data to Create

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, World!")
	// })

	// http.HandleFunc("/mani", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Manikanta!")
	// })

	http.HandleFunc("/Createdata", handleRequest)

	//http.HandleFunc("/generate", handleRequest)

	http.ListenAndServe(":8080", nil)

}

func userRequest(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	var name, age, location string = request.FormValue("name"), request.FormValue("age"), request.FormValue("location")
	var userData []byte

	writer.Header().Set("Content-Type", "application/json")

	if name == "" {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			"Could not determine the desired name.",
		)
		return
	}

	ageNumber, err := strconv.Atoi("age")
	if err != nil || age == "" {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode("Could not determine the Person age.")
		return
	}

	if location == "" {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			"Could not determine the desired location.",
		)
		return
	}

	userCode := User{Name: name, Age: ageNumber, Location: location}
	userData, err = userCode.CreateUser()
	if err != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			fmt.Sprintf("Could not create User . %v", err),
		)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(userData)
}

func (user *User) CreateUser() ([]byte, error) {
	userCode, err := User.Encode(user.Name, user.Age, user.Location)
	if err != nil {
		return nil, fmt.Errorf("could not generate a User code: %v", err)
	}
	return userCode, nil
}

func (code *simpleQRCode) Generate() ([]byte, error) {
	qrCode, err := qrcode.Encode(code.Content, qrcode.Medium, code.Size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	return qrCode, nil
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	var size, content string = request.FormValue("size"), request.FormValue("content")
	var codeData []byte

	writer.Header().Set("Content-Type", "application/json")

	if content == "" {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			"Could not determine the desired QR code content.",
		)
		return
	}

	qrCodeSize, err := strconv.Atoi(size)
	if err != nil || size == "" {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode("Could not determine the desired QR code size.")
		return
	}

	qrCode := simpleQRCode{Content: content, Size: qrCodeSize}
	codeData, err = qrCode.Generate()
	if err != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			fmt.Sprintf("Could not generate QR code. %v", err),
		)
		return
	}

	writer.Header().Set("Content-Type", "image/png")
	writer.Write(codeData)
}

type simpleQRCode struct {
	Content string
	Size    int
}
