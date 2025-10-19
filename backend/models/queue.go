package models

type Queue struct {
	ID uint `json:"id"`
	DoctorID uint `json:"doctor_id"`
	Doctor *Doctor `json:"doctor" gorm:"foreignKey:DoctorID;references:ID"`
	PatientID uint `json:"patient_id"`
	Patient *Patient `json:"patient" gorm:"foreignKey:PatientID;references:ID"`
	TicketID  uint     `json:"ticket_id"`
	Ticket    *Ticket  `json:"ticket" gorm:"foreignKey:TicketID;references:ID"`
	Position int `json:"position"`
}