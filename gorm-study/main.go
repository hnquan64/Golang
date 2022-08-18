package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// Connection to MySQL
	DB_USER := "root"
	DB_PASS := "hnq641999"
	DB_NAME := "gormDB"
	DB_HOST := "localhost"
	DB_PORT := "3306"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	// Connecting
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error %s when connect to MySQL", err)
	} else {
		fmt.Println("Connected!")
	}

}
