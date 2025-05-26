package main

import (
	"moviedb/pkg/movies"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(movies.LoggerMiddleware())
	store := movies.NewMemStore()
	moviesController := movies.NewMoviesController(store)
	router.GET("/movies", moviesController.ListMovies)
	router.GET("/movies/:name", moviesController.GetMovie)
	router.DELETE("/movies/:name", moviesController.DeleteMovie)

	authGroup := router.Group("/api")
	authGroup.Use(movies.AuthMiddleware())
	{
		authGroup.POST("/movies", moviesController.CreateMovie)
		authGroup.PUT("/movies/:name", moviesController.UpdateMovie)
	}

	router.Run(":8080")
}
