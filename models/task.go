package models

import (
	"time"

	"gorm.io/gorm"
)

// Constantes para estados de tarea
const (
	StatusPending    = "pending"
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"
	StatusBlocked    = "blocked"
)

// Constantes para prioridades
const (
	PriorityLow    = "low"
	PriorityMedium = "medium"
	PriorityHigh   = "high"
	PriorityUrgent = "urgent"
)

type Task struct {
	gorm.Model
	NameTask    string     `json:"nameTask" gorm:"type:varchar(255);not null"`
	Description string     `json:"description" gorm:"type:text"`
	StatusTask  string     `json:"statusTask" gorm:"type:varchar(50);default:'pending'"`
	Priority    string     `json:"priority" gorm:"type:varchar(50);default:'medium'"`
	DueDate     *time.Time `json:"dueDate"`      // Fecha límite
	StartDate   *time.Time `json:"startDate"`    // Fecha de inicio
	CompletedAt *time.Time `json:"completedAt"`  // Fecha finalización

	// Relaciones
	ProjectID uint    `json:"projectId"`
	Project   Project `gorm:"foreignKey:ProjectID"`

	AssignedToID uint `json:"assignedToId"`
	AssignedTo   User `gorm:"foreignKey:AssignedToID"`

	CreatedByID uint `json:"createdById"`
	CreatedBy   User `gorm:"foreignKey:CreatedByID"`

	Collaborators []User    `gorm:"many2many:task_collaborators;" json:"collaborators"`
	Comments      []Comment `json:"comments"`
}
