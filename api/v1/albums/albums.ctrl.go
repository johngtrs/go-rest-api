package albums

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/johngtrs/go-rest-api/database/models"
)

type AlbumRepository = models.AlbumRepository

func list(c *gin.Context) {
	albumRepository := AlbumRepository{DB: c.MustGet("db").(*sqlx.DB)}
	albums, err := albumRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func readById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please send an int value"})
		return
	}

	albumRepository := AlbumRepository{DB: c.MustGet("db").(*sqlx.DB)}
	alb, err := albumRepository.FindFirst(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alb)
}

func listByArtist(c *gin.Context) {
	artist := c.Param("name")

	albumRepository := AlbumRepository{DB: c.MustGet("db").(*sqlx.DB)}
	alb, err := albumRepository.FindByArtist(artist)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alb)
}

func create(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	albumRepository := AlbumRepository{DB: c.MustGet("db").(*sqlx.DB)}
	albumID, err := albumRepository.AddAlbum(newAlbum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": albumID})
}
