package models

import (
	"time"
)

type Patient struct {
	ID         uint   `json:"id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	BirthDate  time.Time	`json:"birth_date"`
	Phone string `json:"phone_number"`
	Content string `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Ticket *Ticket `json:"ticket" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Queue []Queue `json:"queues" gorm:"foreignKey:PatientID"`
}

