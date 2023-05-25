package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sr-meeting/meeting-service/model"
	"sr-meeting/meeting-service/service"
)

func eventController() {
	router.GET("/event", getEvents)
	router.GET("/event/:id", getEvent)
	router.GET("/event/meet_id/:meet_id", getEventByMeetId)
	router.DELETE("/event/:id", removeEvent)
	router.POST("/event", addEvent)
	router.PUT("/event", updateEvent)
}

func getEvents(c *gin.Context) {
	events, err := service.GetEvents()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	event, err := service.GetEventById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, event)
}

func getEventByMeetId(c *gin.Context) {
	id := c.Param("meet_id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given meet_id is empty"})
		return
	}

	event, err := service.GetEventByMeetId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, event)
}

func removeEvent(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	err := service.RemoveEventById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, "")
}

func addEvent(c *gin.Context) {
	var event model.Event
	if err := c.BindJSON(&event); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.AddEvent(event)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

func updateEvent(c *gin.Context) {
	var event model.Event
	if err := c.BindJSON(&event); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.UpdateEvent(event)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}
