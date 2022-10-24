package main

import (
	"kathotainment/database"

	"github.com/gin-gonic/gin"

	"kathotainment/handlers"
	"kathotainment/migrations"
)

func main() {

	// Initialize the database.
	database.InitDB()
	// Initialize the migrations
	migrations.Migrate()

	router := gin.Default()
	router.GET("/activities", handlers.GetActivities)
	router.PATCH("/activities", handlers.UpdateActivities)
	router.POST("/activities", handlers.PostActivities)
	router.GET("/activities/:id", handlers.GetActivityById)
	router.DELETE("/activities/:id", handlers.DeleteActivity)

	router.Run("localhost:8080")
}
