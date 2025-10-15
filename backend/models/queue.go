package models

type Queue struct {
	ID uint `json:"id"`
	DoctorID uint `json:"doctor_id"`
	Doctor *Doctor `json:"doctor" gorm:"foreignKey:DoctorID;references:ID"`
	PatientID uint `json:"patient_id"`
	Patient *Patient `json:"patient" gorm:"foreignKey:PatientID;references:ID"`
	Position int `json:"position"`
}