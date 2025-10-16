package controllers

import (
	"clinicQueue/config"
	"clinicQueue/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PatientInput struct {
	LastName       string `json:"last_name" binding:"required"`
	FirstName      string `json:"first_name" binding:"required"`
	MiddleName     string `json:"middle_name" binding:"required"`
	BirthDate      string `json:"birth_date" binding:"required"`
	Phone          string `json:"phone_number" binding:"required,min=11,max=11"`
	PassportNumber string `json:"passport_number" binding:"required,min=11,max=11"`
	PolicyOMS      string `json:"policy_oms" binding:"required,min=16,max=16"`
	Content        string `json:"content"`
}

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
	var input PatientInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	var existingPatient models.Patient

	if err := config.DB.Where("phone_number = ?", input.Phone).First(&existingPatient).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Номер телефона уже тспользуется"})
		return
	}

	if err := config.DB.Where("passport_number = ?", input.PassportNumber).First(&existingPatient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Номер паспорта уже используется"})
		return
	}

	if err := config.DB.Where("policy_oms = ?", input.PolicyOMS).First(&existingPatient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Номер полиса уже используется"})
		return
	}

	patient := models.Patient{
		LastName:       input.LastName,
		FirstName:      input.FirstName,
		MiddleName:     input.MiddleName,
		BirthDate:      input.BirthDate,
		Phone:          input.Phone,
		PassportNumber: input.PassportNumber,
		PolicyOMS:      input.PolicyOMS,
		Content:        input.Content,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := config.DB.Create(patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erroro": "Не удалось создать пациента"})
		return
	}

	c.JSON(http.StatusCreated, patient)
}
