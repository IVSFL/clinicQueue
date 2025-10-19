package controllers

import (
	"clinicQueue/config"
	"clinicQueue/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDoctorQueue(c *gin.Context){
	doctorID := c.Param("id")

	var queue []models.Queue
	if err := config.DB.Preload("Patient").Preload("Ticket").Where("doctor_id = ?", doctorID).Order("position ASC").Find(&queue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения очереди"})
		return
	}

	c.JSON(http.StatusOK, queue)
}