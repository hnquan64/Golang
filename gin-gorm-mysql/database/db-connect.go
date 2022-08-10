package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Function for check load data from .env file
// func getEnv() {
// 	err := godotenv.Load("../.env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

// Create a database fetch from info connection
var Db *gorm.DB

func InitDB() *gorm.DB {
	Db = ConnectDB()
	return Db
}

// Connection to mysql database by info in .env file
func ConnectDB() *gorm.DB {

	// getEnv()

	// DB_USER := os.Getenv("DB_USER")
	// DB_PASS := os.Getenv("DB_PASS")
	// DB_NAME := os.Getenv("DB_NAME")
	// DB_HOST := os.Getenv("DB_HOST")
	// DB_PORT := os.Getenv("DB_PORT")

	DB_USER := "root"
	DB_PASS := "hnq641999"
	DB_NAME := "gin_gorm"
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
