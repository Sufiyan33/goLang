package models

import (
	"github.com/Sufiyan33/goLang/go-bookstore-management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Price       int    `json:"price"`
}

// Invoke init methods to connect database
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
