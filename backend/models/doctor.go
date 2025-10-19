package models

type Doctor struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	UserID           uint           `json:"user_id"`
	User             *User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	LastName         string         `json:"last_name"`
	FirstName        string         `json:"first_name"`
	MiddleName       string         `json:"middle_name"`
	SpecializationID uint           `json:"specialization_id"` // добавлено
	Specialization   Specialization `json:"specialization" gorm:"foreignKey:SpecializationID"`
	Office           string         `json:"office"`
	Role             string         `json:"role"`

	Queue []Queue `json:"queue" gorm:"foreignKey:DoctorID"`
}
