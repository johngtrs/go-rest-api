package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	apiv1 "github.com/johngtrs/go-rest-api/api/v1"
)

func BuildRoutes(r *gin.Engine, db *sqlx.DB) {
	api := r.Group("/api")
	apiv1.BuildRoutes(api, db)
}
