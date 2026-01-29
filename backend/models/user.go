package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Doctor *Doctor `json:"doctor" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Admin *Admin `json:"admin" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
