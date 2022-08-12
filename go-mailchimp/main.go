package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	api := os.Getenv("API_KEY")
	fmt.Println(api)
}
