package main

import (
	"clinicQueue/config"
	"clinicQueue/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env", err)
	}

	config.ConnectDatabase()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8000")
}
