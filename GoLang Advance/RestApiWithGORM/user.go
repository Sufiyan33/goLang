package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func initializeDatabse() {
	db, err := sql.Open("mysql", "root:Mysql@@123@tcp(127.0.0.1:3306)/sufiyan")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect databse")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&User{})
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating new user with POST method")

	db, err := sql.Open("mysql", "root:Mysql@@123@tcp(127.0.0.1:3306)/sufiyan")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect databse")
	}
	defer db.Close()

	params := mux.Vars(r)
	name := params["name"]
	email := params["email"]
	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "Record succesfully created")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all user with GET method")

	db, err := sql.Open("mysql", "root:Mysql@@123@tcp(127.0.0.1:3306)/sufiyan")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect databse")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating user details with PUT method")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting user details with DELTE method")
}
