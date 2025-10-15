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
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=6"`
	LastName       string `json:"last_name" binding:"required"`
	FirstName      string `json:"first_name" binding:"required"`
	MiddleName     string `json:"middle_name" binding:"required"`
	Specialization string `json:"specialization" binding:"required"`
	Office         string `json:"office" binding:"required"`
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

func parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(tokent *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}

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
			UserID:         user.ID,
			LastName:       input.LastName,
			FirstName:      input.FirstName,
			MiddleName:     input.MiddleName,
			Specialization: input.Specialization,
			Office:         input.Office,
		}
		if err := tx.Create(&doctor).Error; err != nil {
			return err
		}

		user.Doctor = &doctor

		token, err := getToken(int(user.ID))
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

}
