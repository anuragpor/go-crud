package controllers

import (
	"fmt"

	"github.com/anurag/go-crud/initializers"
	"github.com/anurag/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body

	// Create a post
	post := models.Post{Title: "Hello", Body: "Post Body"}
	result := initializers.DB.Create(&post) // pass pointer of data to Create

	fmt.Println(post.ID, result.Error, result.RowsAffected)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}
