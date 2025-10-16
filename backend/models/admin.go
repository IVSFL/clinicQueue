package models

type Admin struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	UserID     uint   `json:"user_id"`
	User       *User  `json:"user" gorm:"foreignKey:UserID;references:ID"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Role       string `json:"role"`
}
