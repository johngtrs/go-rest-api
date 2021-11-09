package album

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func BuildRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	albums := rg.Group("/albums")

	r := NewAlbumRepository(db)
	s := NewAlbumService(r)
	h := NewAlbumHandler(s)

	albums.GET("/", h.List)
	albums.GET("/:id", h.ReadById)
	albums.GET("/artist/:name", h.ListByArtist)
	albums.POST("/", h.Create)
}
