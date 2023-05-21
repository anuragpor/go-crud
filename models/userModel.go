package models

import (
	"time"
)

type User struct {
	Email       string `gorm:"primaryKey;not null"`
	Name        string `gorm:"not null"`
	DateOfBirth time.Time
	Movies      []*Movie `gorm:"many2many:user_movies;"`
}
