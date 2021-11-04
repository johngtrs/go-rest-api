package movies

import "github.com/gin-gonic/gin"

func BuildRoutes(r *gin.RouterGroup) {
	movies := r.Group("/movies")
	{
		movies.GET("/", list)
	}
}
