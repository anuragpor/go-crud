package models

import (
	"time"

	"gorm.io/gorm"
)

// type User struct {
// 	Email       string `gorm:"primaryKey"`
// 	Name        string `gorm:"size:255"`
// 	DateOfBirth time.Time
// }

// type Movie struct {
// 	MovieID    int    `gorm:"primaryKey"`
// 	MovieTitle string `gorm:"size:255"`
// }

type Post struct {
	gorm.Model
	Title string
	Body  string
}

type User struct {
	Email       string `gorm:"primaryKey;not null"`
	Name        string `gorm:"not null"`
	DateOfBirth time.Time
	Movies      []*Movie `gorm:"many2many:user_movies;"`
}
type Movie struct {
	Movie_id    int     `gorm:"primaryKey;autoIncrement;not null"`
	Movie_title string  `gorm:"not null"`
	Users       []*User `gorm:"many2many:user_movies;"`
}
