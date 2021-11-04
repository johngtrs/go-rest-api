package api

import (
	"github.com/gin-gonic/gin"
	apiv1 "github.com/johngtrs/go-rest-api/api/v1"
)

func BuildRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		apiv1.BuildRoutes(api)
	}
}
