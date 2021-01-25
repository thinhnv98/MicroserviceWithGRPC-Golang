package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load("env/dev.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Instance Server
	server := gin.Default()

	//Routes
	fmt.Println(server)
}
