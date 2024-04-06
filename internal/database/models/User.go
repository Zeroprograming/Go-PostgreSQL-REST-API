package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model

	Id        string `gorm:"primaryKey"`
	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);not null"`
	Password  string
	Tasks     []Task
}

// SetPassword hashea la contraseña del usuario
func (u *User) SetPassword(password string) error {
	// Generar el hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifica si la contraseña proporcionada coincide con la contraseña almacenada del usuario
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
