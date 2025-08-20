package models

import (
	"errors"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name           string `json:"name"`
	Email          string `json:"email" gorm:"unique"`
	Password       string `json:"password"` // Omitido en respuestas JSON "-"
	Age            int    `json:"age"`
	CreatedTasks   []Task `gorm:"foreignKey:CreatedByID"`
	AssignedTasks  []Task `gorm:"foreignKey:AssignedToID"`
	Collaborations []Task `gorm:"many2many:task_collaborators;"`
	Comments       []Comment
}

// Cifra la contraseña antes de guardarla en la base de datos
func (u *User) HashPassword() error {
	if u.Password == "" {
		return errors.New("la contraseña no puede estar vacia")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Verifica si la contraseña coincide con el hash
func (u *User) CheckPassword(password string) error {
	if u.Password == "" {
		return errors.New("contraseña no establecida")
	}
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}