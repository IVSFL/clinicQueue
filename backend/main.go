package main

import (
	"log"
	"clinicQueue/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env", err);
	}
	
	config.ConnectDatabase();

	r := gin.Default();
	r.Run(":8000");
}