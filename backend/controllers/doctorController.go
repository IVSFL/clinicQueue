package controllers

import(
	"net/http"
	"clinicQueue/models"
	"clinicQueue/config"

	"github.com/gin-gonic/gin"
)

func GetAllDoctors(c *gin.Context) {
	var doctors []models.Doctor

	if err := config.DB.Find(&doctors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
		return
	}

	c.JSON(http.StatusOK, doctors)
}

func GetDoctor(c *gin.Context) {
	id := c.Param("id")
	var doctor models.Doctor

	if err := config.DB.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Врач не найден"})
		return
	}

	c.JSON(http.StatusOK, doctor)
}

func CreateDoctor(c *gin.Context){
	var doctor models.Doctor

	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать врача"})
		return
	}

	c.JSON(http.StatusCreated, doctor)
}

func UpdateDoctor(c *gin.Context){
	id := c.Param("id")
	var doctor models.Doctor

	if err := config.DB.Find(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}

	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&doctor)
	c.JSON(http.StatusOK, doctor)
}