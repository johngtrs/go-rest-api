package movie

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func BuildRoutes(r *gin.RouterGroup, db *sqlx.DB) {
	movies := r.Group("/movies")

	movies.GET("/", list)
}
