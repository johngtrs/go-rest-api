package albums

import "github.com/gin-gonic/gin"

func BuildRoutes(r *gin.RouterGroup) {
	albums := r.Group("/albums")
	{
		albums.GET("/", list)
		albums.GET("/:id", readById)
		albums.POST("/", create)
	}
}
