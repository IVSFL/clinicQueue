package controllers

import (
	"clinicQueue/config"
	"clinicQueue/models"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
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

	config.DB.Model(&models.Ticket{}).Where("id = ?", next.TicketID).Update("status", "processed")

	now := time.Now()
	config.DB.Model(&models.Ticket{}).Where("id = ?", next.ID).Update("called_at", &now)

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

	config.DB.Model(&models.Ticket{}).Where("id = ?", queueItem.TicketID).Update("status", "processed")

	now := time.Now()
	config.DB.Model(&models.Ticket{}).Where("id = ?", queueItem.ID).Update("called_at", &now)

	var doctor models.Doctor
	config.DB.First(&doctor, queueItem.DoctorID)
	BroadcastCall(*queueItem.Patient, doctor, queueItem.Ticket.TicketNumber)

	c.JSON(http.StatusOK, gin.H{
		"message":       "Пациент вызван вручную",
		"ticket_number": queueItem.Ticket.TicketNumber,
		"patient":       queueItem.Patient,
	})
}
