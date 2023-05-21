package routes

import (
	"github.com/anurag/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	//user routes
	router.POST("/users", controllers.UserCreate)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:Email", controllers.UserShow)
	router.PUT("/users/:Email", controllers.UserUpdate)
	router.DELETE("/users/:Email", controllers.UserDelete)
	router.POST("users/:Email", controllers.UserAddMovie)
	router.DELETE("user-movie", controllers.DeleteMovie)
	router.GET("user-movie", controllers.ListUserMovies)

	// movies routes
	router.POST("/movies", controllers.MovieCreate)
	router.GET("/movies", controllers.GetAllMovies)
	router.GET("/movies/:id", controllers.GetMovieById)
}
