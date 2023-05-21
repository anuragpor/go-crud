package controllers

import (
	"fmt"

	"github.com/anurag/go-crud/initializers"
	"github.com/anurag/go-crud/models"
	"github.com/gin-gonic/gin"
)

func MovieCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Movie_title string
	}

	c.Bind(&body)
	// Create a user
	movie := models.Movie{Movie_title: body.Movie_title}
	result := initializers.DB.Create(&movie) // pass pointer of data to Create

	if result.Error != nil {
		fmt.Println("Error occured", result.Error)
		c.Status(400)
		return
	}
	// Return it
	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func GetAllMovies(c *gin.Context) {
	// Get the movie
	var movies []models.Movie
	initializers.DB.Find(&movies)

	// Respond with them
	c.JSON(200, gin.H{
		"movies": movies,
	})
}

func GetMovieById(c *gin.Context) {
	//Get id from url
	id := c.Param("id")

	// Get the post
	var movie models.Movie
	initializers.DB.First(&movie, id)

	// Respond with them
	c.JSON(200, gin.H{
		"movie": movie,
	})
}
