package controllers

import(
	"net/http"
	"clinicQueue/models"
	"clinicQueue/config"

	"github.com/gin-gonic/gin"
)

func GetAllSpecialization(c *gin.Context) {
	var specialization []models.Specialization

	if err := config.DB.Find(&specialization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
		return
	}

	c.JSON(http.StatusOK, specialization)
}