package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/johngtrs/go-rest-api/api"
	"github.com/johngtrs/go-rest-api/database"
	"github.com/johngtrs/go-rest-api/middleware"
	"github.com/johngtrs/go-rest-api/validation"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	if os.Getenv("GO_ENV") == "DOCKER-DEV" {
		err = godotenv.Load(".env.docker")
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	validation.InitValidator()

	db, err := database.Initialize()
	if err != nil {
		panic(err)
	}

	app := gin.Default()

	app.Use(middleware.ErrorHandler)

	api.BuildRoutes(app, db)
	app.Run(":" + os.Getenv("GO_LOCAL_PORT"))
}
