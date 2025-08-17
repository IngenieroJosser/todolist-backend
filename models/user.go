package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name  				string `json:"name"`
    Email 				string `json:"email" gorm:"unique"`
		Password      string `json:"password"`
		Age           int    `json:"age"`

    // Relaciones
	CreatedTasks   []Task `gorm:"foreignKey:CreatedByID"`
	AssignedTasks  []Task `gorm:"foreignKey:AssignedToID"`
	Collaborations []Task `gorm:"many2many:task_collaborators;"`
	Comments       []Comment
}
