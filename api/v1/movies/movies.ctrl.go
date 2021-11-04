package movies

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johngtrs/go-rest-api/database/models"
)

func list(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	movies, err := models.GetMovies(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}
