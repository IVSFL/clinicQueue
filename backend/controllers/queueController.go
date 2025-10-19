package controllers

import (
	"clinicQueue/config"
	"clinicQueue/models"
	"net/http"
	"time"

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

func CallNext(c *gin.Context) {
	doctorID := c.Param("id")

	var next models.Queue

	if err := config.DB.Preload("Patient").Preload("Ticket").Where("doctor_id = ?", doctorID).Order("position ASC").First(&next).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Очередь пуста"})
		return
	}

	config.DB.Model(&models.Ticket{}).Where("id = ?", next.TicketID).Update("status", "processed")

	now := time.Now()
	config.DB.Model(&models.Ticket{}).Where("id = ?", next.ID).Update("called_at", &now)

	c.JSON(http.StatusOK, gin.H{
		"message": "Пациент вызван",
		"ticket_number": next.Ticket.TicketNumber,
		"patient": next.Patient,
	})
}

func CallList(c *gin.Context) {
	doctorID := c.Param("id")
	patientID := c.Param("patient_id")

	var queueItem models.Queue
	if err := config.DB.Preload("Patient").Preload("Ticket").Where("doctor_id = ? AND patient_id = ?", doctorID, patientID).First(&queueItem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пациент не найден"})
		return
	}

	config.DB.Model(&models.Ticket{}).Where("id = ?", queueItem.TicketID).Update("status", "processed")

	now := time.Now()
	config.DB.Model(&models.Ticket{}).Where("id = ?", queueItem.ID).Update("called_at", &now)

	c.JSON(http.StatusOK, gin.H{
		"message": "Пациент вызван вручную",
		"ticket_number": queueItem.Ticket.TicketNumber,
		"patient": queueItem.Patient,
	})
}