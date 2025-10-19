package models

type Specialization struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
	Prefix string `json:"prefix" gorm:"size:1"`
}
