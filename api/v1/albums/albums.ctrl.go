package albums

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/johngtrs/go-rest-api/database/models"
)

func list(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	albums, err := models.GetAlbums(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func readById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please send an int value"})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	alb, err := models.GetAlbumByID(db, id)
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

	db := c.MustGet("db").(*sql.DB)
	albumID, err := models.AddAlbum(db, newAlbum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": albumID})
}
