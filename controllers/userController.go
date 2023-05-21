package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/anurag/go-crud/initializers"
	"github.com/anurag/go-crud/models"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Email       string
		Name        string
		DateOfBirth time.Time
	}

	c.Bind(&body)
	// Create a user
	user := models.User{Email: body.Email, Name: body.Name}
	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		fmt.Println("Error occured", result.Error)
		c.Status(400)
		return
	}
	// Return it
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserIndex(c *gin.Context) {
	// Get the users
	var users []models.User
	initializers.DB.Find(&users)

	// Respond with them
	c.JSON(200, gin.H{
		"users": users,
	})
}

func UserShow(c *gin.Context) {
	//Get id from url
	id := c.Param("Email")
	// Get the user
	var user models.User
	initializers.DB.First(&user, "Email = ?", id)

	// Respond with them
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserUpdate(c *gin.Context) {
	// get id from url
	id := c.Param("Email")

	//get data from request body
	var body struct {
		Name        string
		DateOfBirth time.Time
	}

	c.Bind(&body)

	//Find the user to update
	var user models.User
	initializers.DB.First(&user, "Email = ?", id)

	// Update it
	initializers.DB.Model(&user).Updates(models.User{
		Name:        body.Name,
		DateOfBirth: body.DateOfBirth,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserDelete(c *gin.Context) {
	// Get the id from the url
	id := c.Param("Email")

	//Delete the user
	var user models.User
	initializers.DB.Where("Email = ?", id).Delete(&user)

	//Respond
	c.Status(200)
}

// func GetMovieById(c *gin.Context) {
// 	//Get id from url
// 	id := c.Param("id")

// 	// Get the post
// 	var movie models.Movie
// 	initializers.DB.First(&movie, id)

// 	// Respond with them
// 	c.JSON(200, gin.H{
// 		"movie": movie,
// 	})
// }

func UserAddMovie(c *gin.Context) {
	email_id := c.Param("Email")

	//get data from request body
	var body struct {
		Movie_id string
	}

	c.Bind(&body)
	fmt.Println("Found movie Id ", body.Movie_id)
	// Find that movie by movie_id
	var movie models.Movie
	initializers.DB.First(&movie, body.Movie_id)

	fmt.Println("Found movie ", movie.Movie_title)

	//Find the user to update
	var user models.User
	initializers.DB.First(&user, "Email = ?", email_id)

	initializers.DB.Model(&user).Association("Movies").Append(&movie)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteMovie(c *gin.Context) {
	// take the url data
	email_id := c.Query("email")
	movieIDStr := c.Query("id")

	movieIDToDelete, _ := strconv.Atoi(movieIDStr)
	var user models.User

	// Find the required user
	result := initializers.DB.Preload("Movies").First(&user, "Email = ?", email_id)

	if result.Error != nil {
		fmt.Println("Failed to retrieve user", result.Error)
		return
	}
	// Find the movie to delete from the user's movies
	var movieToDelete models.Movie
	for i, movie := range user.Movies {
		if movie.Movie_id == movieIDToDelete {
			movieToDelete = *movie
			// Remove the movie from the user's movies slice
			user.Movies = append(user.Movies[:i], user.Movies[i+1:]...)
			break
		}
	}
	// Remove the movie from the user's movies and the join table
	initializers.DB.Model(&user).Association("Movies").Delete(&movieToDelete)
	c.Status(200)
}
