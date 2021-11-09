package movie

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service MovieService
}

func NewMovieHandler(service MovieService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) List(c *gin.Context) {
	movies, err := h.service.ListService()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (h *Handler) ReadById(c *gin.Context) {
	id := c.Param("id")
	movie, err := h.service.readByIdService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *Handler) Create(c *gin.Context) {
	var newMovie Movie

	if err := c.BindJSON(&newMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	movieID, err := h.service.createService(newMovie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": movieID})
}
