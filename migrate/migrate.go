package main

import (
	"github.com/anurag/go-crud/initializers"
	"github.com/anurag/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
}
