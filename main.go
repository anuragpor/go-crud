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
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.POST("/users", controllers.UserCreate)
	r.GET("/users", controllers.UserIndex)
	r.GET("/users/:Email", controllers.UserShow)
	r.PUT("/users/:Email", controllers.UserUpdate)
	r.DELETE("/users/:Email", controllers.UserDelete)

	r.POST("users/:Email", controllers.UserAddMovie)
	r.DELETE("user-movie", controllers.DeleteMovie)

	r.POST("/movies", controllers.MovieCreate)
	r.GET("/movies", controllers.GetAllMovies)
	r.GET("/movies/:id", controllers.GetMovieById)

	r.Run() // listen and serve on 0.0.0.0:8080
}
