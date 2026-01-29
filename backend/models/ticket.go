package models

import "time"

type Ticket struct {
	ID               uint            `json:"id" gorm:"primaryKey"`
	PatientID        uint            `json:"patient_id"`
	Patient          *Patient        `json:"patient" gorm:"foreignKey:PatientID;references:ID"`
	SpecializationID uint            `json:"specialization_id"`
	Specialization   *Specialization `json:"specialization" gorm:"foreignKey:SpecializationID"`
	TicketNumber     string          `json:"ticket_number"`
	Status           string          `json:"status"`
	CalledAt         time.Time       `json:"called_at"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}
