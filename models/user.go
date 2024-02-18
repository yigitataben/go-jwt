package models

import "time"

// User struct represents a user entity.
type User struct {
	UserID       uint      `json:"user_id" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"created_at"`
	UserName     string    `json:"user_name" gorm:"unique;not null"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	EmailAddress string    `json:"email_address" gorm:"unique;not null"`
	UserPassword string    `json:"user_password" gorm:"not null"`
}
