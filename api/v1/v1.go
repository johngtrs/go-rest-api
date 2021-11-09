package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/johngtrs/go-rest-api/api/v1/album"
	"github.com/johngtrs/go-rest-api/api/v1/movie"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func BuildRoutes(r *gin.RouterGroup, db *sqlx.DB) {
	v1 := r.Group("/v1")

	v1.GET("/ping", ping)
	album.BuildRoutes(v1, db)
	movie.BuildRoutes(v1, db)
}
