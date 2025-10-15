package models

import "time"

type Ticket struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	PatientID    uint      `json:"patient_id"`
	Patient      *Patient  `json:"patient" gorm:"foreignKey:PatientID;references:ID"`
	TicketNumber string    `json:"ticket_number"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
