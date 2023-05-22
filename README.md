
# GoLang Movies Rest-API

API handles the following request

	*Register user accepting email id, name, date of birth
	*List all User and get a user by its email_id.
	*Update User name and date of birth by email.
	*Add movies in the movies table accepting Movies_title
	*Get movie by ID.
	*Add a list of movies the user has watched.
	*Delete a movie from a list of movies of a user.
    	*List all the movies the user has watched.





## Setup Environment Variables
* Setup the your postgres database string in .env file.

* Set the PORT variable.
## Install Dependencies
* go mod tidy

## Start the Server
Use the follwing command to run the server.
* go run migrate/migrate.go  (Auto migrate feature of gorm to create user, movies and Many to Many relationship table user_movies)
* CompileDaemon -command="./go-crud" (Start listening on the PORT)






## Models

* model/movieModel.go has following schema

type Movie struct 
    {

    Movie_id    int     `gorm:"primaryKey;autoIncrement;not null"`
	Movie_title string  `gorm:"not null"`
	Users       []*User `gorm:"many2many:user_movies;"`
    }

* model/userModel.go has following schema

type User struct {
	
    Email       string `gorm:"primaryKey;not null"`
	Name        string `gorm:"not null"`
	DateOfBirth time.Time
	Movies      []*Movie `gorm:"many2many:user_movies;"`
    }










## Authors

- [@anuragporwal](https://github.com/anuragpor)

