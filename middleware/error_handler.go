package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johngtrs/go-rest-api/utils"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

loopErrors:
	for _, err := range c.Errors {
		switch err.Err {
		case utils.ErrBadRequest:
			c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrBadRequest.Error()})
			break loopErrors
		case utils.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{"error": utils.ErrUnauthorized.Error()})
			break loopErrors
		case utils.ErrPaymentRequired:
			c.JSON(http.StatusPaymentRequired, gin.H{"error": utils.ErrPaymentRequired.Error()})
			break loopErrors
		case utils.ErrForbidden:
			c.JSON(http.StatusForbidden, gin.H{"error": utils.ErrForbidden.Error()})
			break loopErrors
		case utils.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": utils.ErrNotFound.Error()})
			break loopErrors
		case utils.ErrMethodNotAllowed:
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": utils.ErrMethodNotAllowed.Error()})
			break loopErrors
		case utils.ErrNotAcceptable:
			c.JSON(http.StatusNotAcceptable, gin.H{"error": utils.ErrNotAcceptable.Error()})
			break loopErrors
		case utils.ErrConflict:
			c.JSON(http.StatusConflict, gin.H{"error": utils.ErrConflict.Error()})
			break loopErrors
		case utils.ErrInternalServerError:
			c.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrInternalServerError.Error()})
			break loopErrors
		default:
			c.JSON(-1, gin.H{"error": err.Err.Error()})
			break loopErrors
		}
	}
}
