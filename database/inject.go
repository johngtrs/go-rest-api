package database

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Inject(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
