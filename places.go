package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type place struct {
	ID string `json:"id`
	Name string `json:"name"`
	Location string `json:"location"`
	Url string `json:"url"`
	Img_path string `json:"image_path"`
}

var places = []place{
	{ID:"1", Name:"Joe's Bakery", Location:"123 Your Mom's Place", Url:"", Img_path:""},
	{ID:"2", Name:"Kathy Krafts", Location:"123 Bubz Place", Url:"", Img_path:""},
}

func getPlaceById(c *gin.Context) {
	id := c.Param("id")
	for _, p := range places {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	} 
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Place not found."})
}

func getPlaces(c *gin.Context){
	c.IndentedJSON(http.StatusOK, places)
}

func postPlaces(c *gin.Context){
	var newPlace place

	// Call BindJSON to bind the received JSON to
	// newPlace.
	if err := c.BindJSON(&newPlace); err != nil {
		return
	}

	// Add the new place to the slice of places
	places = append(places, newPlace)
	c.IndentedJSON(http.StatusCreated, newPlace)
}

func deletePlaces(c *gin.Context){
	c.IndentedJSON(http.StatusOK, places)
}

func updatePlaces(c *gin.Context){
	c.IndentedJSON(http.StatusOK, places)
}
