package main

import (
	"github.com/anurag/go-crud/initializers"
	"github.com/anurag/go-crud/routes"
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run()
}
