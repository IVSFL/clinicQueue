package controllers

import (
	"clinicQueue/config"
	"clinicQueue/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type RegisterDoctorInput struct {
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,min=6"`
	LastName         string `json:"last_name" binding:"required"`
	FirstName        string `json:"first_name" binding:"required"`
	MiddleName       string `json:"middle_name" binding:"required"`
	SpecializationID uint   `json:"specialization_id" binding:"required"`
	Office           string `json:"office"`
}

type RegisterAdminInput struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	LastName   string `json:"last_name" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	MiddleName string `json:"middle_name" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func getToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// func parseToken(tokenString string) (*jwt.Token, error) {
// 	return jwt.Parse(tokenString, func(tokent *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	})
// }

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterDoctor(c *gin.Context) {
	var input RegisterDoctorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email уже используется"})
		return
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		hashedPassword, err := hashPassword(input.Password)
		if err != nil {
			return err
		}

		user := models.User{
			Email:     input.Email,
			Password:  hashedPassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		doctor := models.Doctor{
			UserID:           user.ID,
			LastName:         input.LastName,
			FirstName:        input.FirstName,
			MiddleName:       input.MiddleName,
			SpecializationID: input.SpecializationID,
			Role:             "doctor",
		}
		if err := tx.Create(&doctor).Error; err != nil {
			return err
		}

		user.Doctor = &doctor

		token, err := getToken(int(user.ID))
		if err != nil {
			return err
		}

		c.JSON(http.StatusCreated, gin.H{
			"user":  user,
			"token": token,
		})
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать врача"})
	}

}

func RegisterAdmin(c *gin.Context) {
	var input RegisterAdminInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email уже используется"})
		return
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		hashedPassword, err := hashPassword(input.Password)
		if err != nil {
			return err
		}

		user := models.User{
			Email:     input.Email,
			Password:  hashedPassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		admin := models.Admin{
			UserID:     user.ID,
			LastName:   input.LastName,
			FirstName:  input.FirstName,
			MiddleName: input.MiddleName,
			Role:       "admin",
		}
		if err := tx.Create(&admin).Error; err != nil {
			return err
		}

		user.Admin = &admin

		token, err := getToken(int(user.ID))
		if err != nil {
			return err
		}

		c.JSON(http.StatusCreated, gin.H{
			"user":  user,
			"token": token,
		})
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать администратора"})
	}
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Preload("Doctor").Preload("Admin").Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Такого Email не существует"})
		return
	}

	if !checkPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "пароль введен не верно"})
		return
	}

	token, err := getToken(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка выдачи токена"})
		return
	}

	if user.Doctor != nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"role":  "doctor",
			"user": gin.H{
				"id":             user.ID,
				"role_id":        user.Doctor.ID,
				"email":          user.Email,
				"first_name":     user.Doctor.FirstName,
				"last_name":      user.Doctor.LastName,
				"middle_name":    user.Doctor.MiddleName,
				"specialization": user.Doctor.Specialization,
				"office":         user.Doctor.Office,
			},
		})
	} else if user.Admin != nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"role":  "admin",
			"user": gin.H{
				"id":          user.ID,
				"role_id":     user.Admin.ID,
				"email":       user.Email,
				"first_name":  user.Admin.FirstName,
				"last_name":   user.Admin.LastName,
				"middle_name": user.Admin.MiddleName,
			},
		})
	}
}
