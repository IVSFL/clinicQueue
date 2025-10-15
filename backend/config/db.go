package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"clinicQueue/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{});
	
	if err != nil{
		log.Fatal("Failed connect to database!", err);
	}

	DB = database;

	fmt.Printf("Connect sucsses");

	err = DB.AutoMigrate(
		&models.User{},
		&models.Doctor{},
		&models.Admin{},
		&models.Patient{},
		&models.Queue{},
		&models.Ticket{},
	)
	if err != nil{
		log.Fatal("Failed to migration", err)
	}

	fmt.Printf("Database migrated sucsses");
}
