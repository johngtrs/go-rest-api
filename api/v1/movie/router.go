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
	movies.POST("/", h.Create)
}
