package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}

var user []User

func main() {

	dsn := "root:ap02BL1426*@tcp(127.0.0.1:3306)/qrcode?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err := sql.Open("mysql", "root:ap02BL1426*@tcp(127.0.0.1:3306)/qrcode")
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	// Check if the database connection is successful
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	router := mux.NewRouter()

	db.AutoMigrate(&User{})

	//user := User{Name: "Chinnodu", Age: 18, Location: "Hyderabad"}

	//db.Create(&user) // pass pointer of data to Create

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, World!")
	// })

	// http.HandleFunc("/mani", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Manikanta!")
	// })

	//http.HandleFunc("/Createdata", handleRequest)
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")

	//http.HandleFunc("/generate", handleRequest)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func getUsers(w http.ResponseWriter, r *http.Request) {

	dsn := "root:ap02BL1426*@tcp(127.0.0.1:3306)/qrcode?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err := sql.Open("mysql", "root:ap02BL1426*@tcp(127.0.0.1:3306)/qrcode")
	if err != nil {
		log.Fatal(err)
	}

	var user []User
	rows := db.Raw("SELECT id, name, age, location FROM users").Scan(&user)
	if rows.Error != nil {
		fmt.Println(">>>>>>>>>>>>>>>>>>>Error", rows.Error)
	}
	fmt.Println("rows >>>>>", user)

	structString := fmt.Sprintf("%+v\n", user)
	w.Write([]byte(structString))

}

func createUser(w http.ResponseWriter, r *http.Request) {

	dsn := "root:ap02BL1426*@tcp(127.0.0.1:3306)/qrcode?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	SQL := "INSERT INTO users (name, age, location) values(?,?,?)"
	details := User{Name: "Mmm", Age: 33, Location: "HYd"}
	db.Exec(SQL, details.Name, details.Age, details.Location)
	fmt.Println(">>>", SQL)

}
