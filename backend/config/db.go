package config

import (
	"fmt"
	"log"
	"os"

	"github.com/glebarez/sqlite" // pure-Go SQLite
	"gorm.io/gorm"

	"clinicQueue/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Для SQLite используем путь к файлу базы данных
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		// Значение по умолчанию
		dbPath = "clinic_queue.db"
	}

	// Создаем директорию если нужно
	err := os.MkdirAll("./data", 0755)
	if err != nil {
		log.Fatal("Failed to create data directory:", err)
	}

	// Формируем полный путь
	fullPath := "./data/" + dbPath

	dsn := fullPath
	database, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to database!", err)
	}

	DB = database

	fmt.Println("Connect success to SQLite database")

	// Настраиваем внешние ключи для SQLite
	DB.Exec("PRAGMA foreign_keys = ON")

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

	fmt.Println("Database migrated successfully")

	// Заполняем начальные данные
	SeedSpecialization()
}

// SeedSpecialization функция остается без изменений
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

		if err := DB.Select("Name", "Prefix").Create(&spec).Error; err != nil {
			log.Println("Ошибка вставки:", err)
		} else {
			log.Println("Специализация добавлена:", name, "=>", prefix)
		}
	}
}