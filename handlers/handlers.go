package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"kathotainment/database"
	"kathotainment/models"
)

func GetActivityById(c *gin.Context) {
	id := c.Param("id")
	var activities models.Activity
	result := database.DBCon.First(&activities, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Activity with ID: " + id + " not found."})
	} else {
		c.IndentedJSON(http.StatusFound, activities)
	}
}

func GetActivities(c *gin.Context) {
	var activities []models.Activity
	result := database.DBCon.Find(&activities)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch activities from database."})
	} else {
		c.IndentedJSON(http.StatusOK, activities)
	}
}

func PostActivities(c *gin.Context) {
	// TODO: Add test for create to database
	var newActivity models.Activity
	// Call BindJSON to bind the received JSON to
	// newActivity.
	if err := c.BindJSON(&newActivity); err != nil {
		return
	}

	// Add the new place to the slice of places
	result := database.DBCon.Create(&newActivity)
	if result.Error != nil {
		panic("Could not add activity to database.")
	}
	c.Status(http.StatusCreated)
}

func DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	var activities models.Activity
	result := database.DBCon.Delete(&activities, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Could not delete activity with ID: " + id + " from database."})
	}
	c.Status(http.StatusNoContent)
}

func UpdateActivities(c *gin.Context) {
	var newActivity models.Activity
	if err := c.BindJSON(&newActivity); err != nil {
		return
	}
	result := database.DBCon.Updates(&newActivity)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Unable to update record."})
	}
	c.Status(http.StatusNoContent)
}
