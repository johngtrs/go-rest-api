package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johngtrs/go-rest-api/httperror"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	responseErrors := map[error]int{
		httperror.ErrBadRequest:          http.StatusBadRequest,          // 400
		httperror.ErrUnauthorized:        http.StatusUnauthorized,        // 401
		httperror.ErrPaymentRequired:     http.StatusPaymentRequired,     // 402
		httperror.ErrForbidden:           http.StatusForbidden,           // 403
		httperror.ErrNotFound:            http.StatusNotFound,            // 404
		httperror.ErrMethodNotAllowed:    http.StatusMethodNotAllowed,    // 405
		httperror.ErrNotAcceptable:       http.StatusNotAcceptable,       // 406
		httperror.ErrConflict:            http.StatusConflict,            // 409
		httperror.ErrInternalServerError: http.StatusInternalServerError, // 500
	}

	for _, err := range c.Errors {
		// Check if the current error exists in the responseErrors array
		if code, ok := responseErrors[err.Err]; ok {
			c.JSON(code, gin.H{"error": err.Err.Error()})
			return
		}
	}
}
