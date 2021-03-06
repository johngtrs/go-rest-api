package movie

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/johngtrs/go-rest-api/model"
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
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) ReadById(c *gin.Context) {
	id := c.Param("id")

	movie, err := h.service.readByIdService(id)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *Handler) Top100(c *gin.Context) {
	movies, err := h.service.Top100Service()
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) Top100Year(c *gin.Context) {
	year := c.Param("year")

	movies, err := h.service.Top100YearService(year)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) MostRented(c *gin.Context) {
	movie, err := h.service.MostRentedService()
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *Handler) MostRentedYear(c *gin.Context) {
	year := c.Param("year")

	movie, err := h.service.MostRentedYearService(year)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *Handler) BestAuthor(c *gin.Context) {
	author, err := h.service.BestAuthorService()
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *Handler) SearchTitle(c *gin.Context) {
	title := c.Param("title")

	movies, err := h.service.SearchByTitleService(title)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) Create(c *gin.Context) {
	var newMovie model.Movie

	if err := c.ShouldBindJSON(&newMovie); err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	movieID, err := h.service.createService(newMovie)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": movieID})
}

func (h *Handler) IncrementRentedNumber(c *gin.Context) {
	var movie model.Movie

	if err := json.NewDecoder(c.Request.Body).Decode(&movie); err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	err := h.service.IncrementRentedNumberService(movie.Title, strconv.Itoa(int(movie.Year)))
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rented number updated"})
}
