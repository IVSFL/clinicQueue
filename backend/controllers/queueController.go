package controllers

import (
	"clinicQueue/config"
	"clinicQueue/models"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex

func WSHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func BroadcastCall(patient models.Patient, doctor models.Doctor, ticketNumber string) {
	mu.Lock()
	defer mu.Unlock()

	message := map[string]interface{}{
		"patient":      patient,
		"doctor":       doctor,
		"ticketNumber": ticketNumber,
		"office":       doctor.Office,
	}

	for client := range clients {
		err := client.WriteJSON(message)
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}

func GetDoctorQueue(c *gin.Context) {
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

	// Меняем статус на "completed" вместо "processed"
	config.DB.Model(&models.Ticket{}).Where("id = ?", next.TicketID).Update("status", "completed")

	now := time.Now()
	config.DB.Model(&models.Ticket{}).Where("id = ?", next.TicketID).Update("called_at", &now)

	config.DB.Delete(&next)

	var doctor models.Doctor
	config.DB.First(&doctor, next.DoctorID)
	BroadcastCall(*next.Patient, doctor, next.Ticket.TicketNumber)

	c.JSON(http.StatusOK, gin.H{
		"message":       "Пациент вызван",
		"ticket_number": next.Ticket.TicketNumber,
		"patient":       next.Patient,
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

	// Меняем статус на "completed" вместо "processed"
	config.DB.Model(&models.Ticket{}).Where("id = ?", queueItem.TicketID).Update("status", "completed")

	now := time.Now()
	config.DB.Model(&models.Ticket{}).Where("id = ?", queueItem.TicketID).Update("called_at", &now)

	config.DB.Delete(&queueItem)

	var doctor models.Doctor
	config.DB.First(&doctor, queueItem.DoctorID)
	BroadcastCall(*queueItem.Patient, doctor, queueItem.Ticket.TicketNumber)

	c.JSON(http.StatusOK, gin.H{
		"message":       "Пациент вызван вручную",
		"ticket_number": queueItem.Ticket.TicketNumber,
		"patient":       queueItem.Patient,
	})
}

func CompletePatient(c *gin.Context) {
	ticketNumber := c.Param("ticket_number")

	var ticket models.Ticket
	if err := config.DB.Where("ticket_number = ?", ticketNumber).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Талон не найден"})
		return
	}

	// Меняем статус на "completed"
	config.DB.Model(&models.Ticket{}).Where("id = ?", ticket.ID).Update("status", "completed")
	now := time.Now()
	config.DB.Model(&models.Ticket{}).Where("id = ?", ticket.ID).Update("called_at", &now)

	c.JSON(http.StatusOK, gin.H{
		"message":       "Прием пациента завершен",
		"ticket_number": ticketNumber,
	})
}

func DeferPatient(c *gin.Context) {
	ticketNumber := c.Param("ticket_number")

	var ticket models.Ticket
	if err := config.DB.Where("ticket_number = ?", ticketNumber).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Талон не найден"})
		return
	}

	config.DB.Model(&models.Ticket{}).Where("id = ?", ticket.ID).Update("status", "deferred")

	c.JSON(http.StatusOK, gin.H{
		"message":       "Пациент отложен",
		"ticket_number": ticketNumber,
	})
}

func DeferPatientForTicket(c *gin.Context) {
	ticketNumber := c.Param("ticket_number")

	var ticket models.Ticket
	if err := config.DB.Where("ticket_number = ?", ticketNumber).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Талон не найден"})
		return
	}

	config.DB.Model(&models.Ticket{}).Where("id = ?", ticket.ID).Update("status", "deferred")

	c.JSON(http.StatusOK, gin.H{
		"message":       "Пациент отложен",
		"ticket_number": ticketNumber,
	})
}

func GetDoctorDeferredQueue(c *gin.Context) {
	doctorID := c.Param("id")

	var doctor models.Doctor
	if err := config.DB.Preload("Specialization").Where("id = ?", doctorID).First(&doctor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}

	var deferredTickets []models.Ticket
	if err := config.DB.
		Preload("Patient").
		Preload("Specialization").
		Where("specialization_id = ? AND status = ?", doctor.SpecializationID, "deferred").
		Order("created_at ASC").
		Find(&deferredTickets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения отложенных пациентов: " + err.Error()})
		return
	}

	var result []map[string]interface{}
	for _, ticket := range deferredTickets {
		result = append(result, map[string]interface{}{
			"ticket_number":  ticket.TicketNumber,
			"patient":        ticket.Patient,
			"specialization": ticket.Specialization,
			"created_at":     ticket.CreatedAt,
			"called_at":      ticket.CalledAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"doctor":         doctor,
		"specialization": doctor.Specialization,
		"deferred":       result,
		"count":          len(result),
	})
}

func CallDeferredPatient(c *gin.Context) {
	doctorID := c.Param("id")
	patientID := c.Param("patient_id")

	// Сначала проверяем существование врача и получаем его специализацию
	var doctor models.Doctor
	if err := config.DB.Preload("Specialization").Where("id = ?", doctorID).First(&doctor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}

	// Ищем отложенный талон для этого пациента с той же специализацией, что у врача
	var ticket models.Ticket
	if err := config.DB.Preload("Patient").Preload("Specialization").
		Where("patient_id = ? AND status = ? AND specialization_id = ?",
			patientID, "deferred", doctor.SpecializationID).
		First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Отложенный талон не найден",
			"details": "У пациента нет отложенного талона для вашей специализации",
		})
		return
	}

	// Обновляем статус талона
	config.DB.Model(&models.Ticket{}).Where("id = ?", ticket.ID).Update("status", "processed")
	now := time.Now()
	config.DB.Model(&models.Ticket{}).Where("id = ?", ticket.ID).Update("called_at", &now)

	BroadcastCall(*ticket.Patient, doctor, ticket.TicketNumber)

	c.JSON(http.StatusOK, gin.H{
		"message":       "Отложенный пациент вызван",
		"ticket_number": ticket.TicketNumber,
		"patient":       ticket.Patient,
		"office":        doctor.Office,
	})
}

func TransferPatient(c *gin.Context) {
	ticketNumber := c.Param("ticket_number")

	var request struct {
		SpecializationID uint `json:"specialization_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ticket models.Ticket
	if err := config.DB.Preload("Patient").
		Where("ticket_number = ?", ticketNumber).
		First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Талон не найден"})
		return
	}

	var doctor models.Doctor
	if err := config.DB.Where("specialization_id = ?", request.SpecializationID).First(&doctor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Врач с указанной специализацией не найден"})
		return
	}

	var maxPosition int
	config.DB.Model(&models.Queue{}).
		Where("doctor_id = ?", doctor.ID).
		Select("COALESCE(MAX(position), 0)").
		Scan(&maxPosition)

	newQueueItem := models.Queue{
		DoctorID:  doctor.ID,
		PatientID: ticket.PatientID,
		TicketID:  ticket.ID,
		Position:  maxPosition + 1,
	}

	if err := config.DB.Create(&newQueueItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении в очередь: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Пациент успешно передан врачу с нужной специализацией",
		"ticket_number": ticket.TicketNumber,
		"doctor":        doctor,
	})
}
