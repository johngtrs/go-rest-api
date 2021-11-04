package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johngtrs/go-rest-api/api/v1/albums"
	"github.com/johngtrs/go-rest-api/api/v1/movies"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func BuildRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", ping)
		albums.BuildRoutes(v1)
		movies.BuildRoutes(v1)
	}
}
