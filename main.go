package main

import (
	"github.com/anurag/go-crud/controllers"
	"github.com/anurag/go-crud/initializers"
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
