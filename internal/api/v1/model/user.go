package model

import "gorm.io/gorm"

// User represents a user in the system.
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);uniqueIndex"`
	Password string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null"`
}

// NewUser creates a new instance of User.
func NewUser(username, password, email string) *User {
	return &User{
		Username: username,
		Password: password,
		Email:    email,
	}
}
