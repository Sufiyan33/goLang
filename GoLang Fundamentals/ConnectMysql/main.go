package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Welcome to connect mySql with go")

	db, err := sql.Open("mysql", "root:Mysql@@123@tcp(127.0.0.1:3306)/sufiyan")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query("INSERT into demos VALUES('GO')")
	if err != nil {
		panic(err)
	}

	defer insert.Close()
	fmt.Println("Successfully inserted data into mysql database")
}
