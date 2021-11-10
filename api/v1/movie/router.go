package movie

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func BuildRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	movies := rg.Group("/movies")

	r := NewMovieRepository(db)
	s := NewMovieService(r)
	h := NewMovieHandler(s)

	movies.GET("/", h.List)
	movies.GET("/:id", h.ReadById)
	movies.GET("/top100", h.Top100)
	movies.GET("/top100/:year", h.Top100Year)
	movies.GET("/most_rented", h.MostRented)
	movies.GET("/most_rented/:year", h.MostRentedYear)
	movies.GET("/best_author", h.BestAuthor)
	movies.GET("/search/:title", h.SearchTitle)
	movies.POST("/", h.Create)
	movies.PATCH("/increment_rented", h.IncrementRentedNumber)
}
