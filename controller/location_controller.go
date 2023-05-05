package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sr-meeting/meeting-service/model"
	"sr-meeting/meeting-service/service"
)

func locationController() {
	router.GET("/location", getLocations)
	router.GET("/location/:id", getLocation)
	router.POST("/location", addLocation)
	router.PUT("/location", updateLocation)
	router.DELETE("/location/:id", removeLocation)
}

func getLocations(c *gin.Context) {
	locations, err := service.GetLocations()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, locations)
}

func getLocation(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	location, err := service.GetLocationById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, location)
}

func removeLocation(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}

func addLocation(c *gin.Context) {
	var location model.Location
	if err := c.BindJSON(&location); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.AddLocation(location)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

func updateLocation(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}
