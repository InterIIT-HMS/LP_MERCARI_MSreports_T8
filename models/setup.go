package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "scar:passloll@tcp(10.46.144.2:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
