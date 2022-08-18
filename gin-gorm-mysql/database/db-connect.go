package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Create a database fetch from info connection
var Db *gorm.DB

func InitDB() *gorm.DB {
	Db = ConnectDB()
	return Db
}

// Connection to mysql database by info in .env file
func ConnectDB() *gorm.DB {

	DB_USER := "root"
	DB_PASS := "*******"
	DB_NAME := "******"
	DB_HOST := "localhost"
	DB_PORT := "3306"

	var err error
	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?parseTime=true&loc=Local"

	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connect to database")
	}

	return db
}
