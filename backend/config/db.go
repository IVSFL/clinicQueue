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

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to database!", err)
	}

	DB = database

	fmt.Printf("Connect sucsses")

	err = DB.AutoMigrate(
		&models.User{},
		&models.Doctor{},
		&models.Admin{},
		&models.Patient{},
		&models.Queue{},
		&models.Ticket{},
		&models.Specialization{},
	)
	if err != nil {
		log.Fatal("Failed to migration", err)
	}

	fmt.Printf("Database migrated success")
}

func SeedSpecialization() {
	specializations := map[string]string{
		"Терапевт":      "А",
		"Хирург":        "Б",
		"Отоларинголог": "В",
		"Невролог":      "Г",
	}

	for name, prefix := range specializations {
		var existing models.Specialization

		// Ищем существующую запись
		result := DB.Where("name = ?", name).First(&existing)

		// Если есть ошибка, не связанная с отсутствием записи
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			log.Println("Ошибка при проверке специализации:", result.Error)
			continue
		}

		// Если запись уже существует — обновляем префикс, если он NULL
		if result.Error == nil {
			if existing.Prefix == "" || existing.Prefix == "NULL" {
				existing.Prefix = prefix
				if err := DB.Model(&existing).Update("Prefix", prefix).Error; err != nil {
					log.Println("Ошибка обновления Prefix:", err)
				} else {
					log.Println("Обновлён Prefix для:", name)
				}
			}
			continue
		}

		// Создаём новую запись, явно указываем поля
		spec := models.Specialization{
			Name:   name,
			Prefix: prefix,
		}

		if err := DB.Debug().Select("Name", "Prefix").Create(&spec).Error; err != nil {
			log.Println("Ошибка вставки:", err)
		} else {
			log.Println("Специализация добавлена:", name, "=>", prefix)
		}
	}
}
