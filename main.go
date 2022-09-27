package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
	router.GET("/places", getPlaces)
	router.GET("/places/:id", getPlaceById)	
	router.POST("/places", postPlaces)

    router.Run("localhost:8080")
}