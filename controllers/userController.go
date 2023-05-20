package controllers

import (
	"fmt"
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
