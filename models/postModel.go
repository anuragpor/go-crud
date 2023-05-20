package models

import "gorm.io/gorm"

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
