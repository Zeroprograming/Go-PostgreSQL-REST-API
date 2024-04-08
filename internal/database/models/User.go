package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User struct represents a user in the system.
type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"` // ID is the unique identifier for the user.
	FirstName string    `gorm:"type:varchar(100);not null"`          // FirstName represents the first name of the user.
	LastName  string    `gorm:"type:varchar(100);not null"`          // LastName represents the last name of the user.
	Email     string    `gorm:"type:varchar(100);not null"`          // Email is the email address of the user.
	Password  string    // Password stores the hashed password of the user.
	Tasks     []Task    // Tasks is a list of tasks associated with the user.
}

// BeforeCreate is a hook that is executed before creating a new record in the database.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	// Generate the hash of the password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}

// SetPassword hashes the user's password.
func (u *User) SetPassword(password string) error {
	// Generate the hash of the password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword checks if the provided password matches the stored hashed password of the user.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
