package movie

import (
	"net/http"
	"strconv"

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

func (h *Handler) Top100(c *gin.Context) {
	movies, err := h.service.Top100Service()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) Top100Year(c *gin.Context) {
	year := c.Param("year")
	movies, err := h.service.Top100YearService(year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) MostRented(c *gin.Context) {
	movie, err := h.service.MostRentedService()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *Handler) MostRentedYear(c *gin.Context) {
	year := c.Param("year")
	movie, err := h.service.MostRentedYearService(year)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *Handler) BestAuthor(c *gin.Context) {
	movie, err := h.service.BestAuthorService()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *Handler) SearchTitle(c *gin.Context) {
	title := c.Param("title")
	movies, err := h.service.SearchByTitleService(title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
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

func (h *Handler) IncrementRentedNumber(c *gin.Context) {
	var movie Movie

	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := h.service.IncrementRentedNumberService(movie.Title, strconv.Itoa(int(movie.Year)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Rented number updated"})
}
