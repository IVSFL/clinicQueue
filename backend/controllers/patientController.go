package controllers

import(
	"net/http"
	"clinicQueue/models"
	"clinicQueue/config"

	"github.com/gin-gonic/gin"
)

func GetAllPatients(c *gin.Context) {
	var patients []models.Patient

	if err := config.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"}) 
		return
	}

	c.JSON(http.StatusOK, patients)
}

func GetPatient(c *gin.Context) {
	id := c.Param("id")

	var patient models.Patient
	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пациент не найден"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func CreatePatient(c *gin.Context) {
	var patient models.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пациента"})
		return
	}

	c.JSON(http.StatusCreated, patient)
}