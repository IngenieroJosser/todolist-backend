package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content"`

	// Relaciones
	TaskID uint `json:"taskId"`
	UserID uint `json:"userId"`

	Task Task `gorm:"foreignKey:TaskID"`
	User User `gorm:"foreignKey:UserID"`
}
