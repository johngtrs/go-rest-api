package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johngtrs/go-rest-api/model"
)

type Handler struct {
	service AlbumService
}

func NewAlbumHandler(service AlbumService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) List(c *gin.Context) {
	albums, err := h.service.ListService()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func (h *Handler) ReadById(c *gin.Context) {
	id := c.Param("id")
	album, err := h.service.readByIdService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

func (h *Handler) ListByArtist(c *gin.Context) {
	artist := c.Param("name")
	albums, err := h.service.listByArtistService(artist)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func (h *Handler) Create(c *gin.Context) {
	var newAlbum model.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	albumID, err := h.service.createService(newAlbum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": albumID})
}
