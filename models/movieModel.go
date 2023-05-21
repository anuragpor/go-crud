package models

type Movie struct {
	Movie_id    int     `gorm:"primaryKey;autoIncrement;not null"`
	Movie_title string  `gorm:"not null"`
	Users       []*User `gorm:"many2many:user_movies;"`
}
