package controllers

import (
	"clinicQueue/config"
	"clinicQueue/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Генерация номера талона по specializationID
func GenerateTicketNumber(specializationID uint) (string, error) {
	var spec models.Specialization
	if err := config.DB.First(&spec, specializationID).Error; err != nil {
		return "", err
	}

	today := time.Now().Format("2006-01-02")
	var count int64
	err := config.DB.Model(&models.Ticket{}).
		Where("specialization_id = ? AND DATE(created_at) = ?", specializationID, today).
		Count(&count).Error
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%03d", spec.Prefix, count+1), nil
}

func CreateTicket(c *gin.Context) {
	var req struct {
		PatientID      uint   `json:"patient_id"`
		Specialization string `json:"specialization"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	// Специализация по названию
	var spec models.Specialization
	if err := config.DB.Where("name = ?", req.Specialization).First(&spec).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Специализация не найдена"})
		return
	}

	//Генерируем номер талона
	ticketNumber, err := GenerateTicketNumber(spec.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации талона"})
		return
	}

	// Создаём талон
	ticket := models.Ticket{
		PatientID:        req.PatientID,
		SpecializationID: spec.ID,
		TicketNumber:     ticketNumber,
		Status:           "waiting",
		CreatedAt:        time.Now(),
	}

	if err := config.DB.Create(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения талона"})
		return
	}

	// Находим врача по специализации
	var doctor models.Doctor
	if err := config.DB.Where("specialization_id = ?", spec.ID).First(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Врач по этой специализации не найден"})
		return
	}

	// Определяем позицию в очереди
	var count int64
	config.DB.Model(&models.Queue{}).Where("doctor_id = ?", doctor.ID).Count(&count)

	queue := models.Queue{
		DoctorID:  doctor.ID,
		PatientID: req.PatientID,
		TicketID:  ticket.ID,
		Position:  int(count + 1),
	}

	if err := config.DB.Create(&queue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления в очередь"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Талон успешно создан и добавлен в очередь",
		"ticket_number":  ticket.TicketNumber,
		"doctor_id":      doctor.ID,
		"patient_id":     req.PatientID,
		"queue_position": queue.Position,
	})
}
