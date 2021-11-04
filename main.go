package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/johngtrs/go-rest-api/api"
	"github.com/johngtrs/go-rest-api/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.Initialize()
	if err != nil {
		panic(err)
	}

	app := gin.Default()
	app.Use(database.Inject(db))
	api.BuildRoutes(app)
	app.Run(":" + os.Getenv("GO_LOCAL_PORT"))
}
