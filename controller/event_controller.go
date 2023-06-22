package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swimresults/meeting-service/dto"
	"github.com/swimresults/meeting-service/model"
	"github.com/swimresults/meeting-service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
)

func eventController() {
	router.GET("/event", getEvents)
	router.GET("/event/:id", getEvent)
	router.GET("/event/meet/:meet_id", getEventsByMeetId)
	router.GET("/event/meet/:meet_id/parts", getEventsAsPartsByMeetId)
	router.GET("/event/meet/:meet_id/event/:event_id", getEventByMeetingAndNumber)

	router.POST("/event", addEvent)
	router.POST("/event/import", importEvent)

	router.DELETE("/event/:id", removeEvent)
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

func getEventsByMeetId(c *gin.Context) {
	id := c.Param("meet_id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given meet_id is empty"})
		return
	}
	events, err := service.GetEventsByMeetId(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, events)
}

func getEventsAsPartsByMeetId(c *gin.Context) {
	id := c.Param("meet_id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given meet_id is empty"})
		return
	}
	parts, err := service.GetEventsAsPartsByMeetId(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, parts)
}

func getEventByMeetingAndNumber(c *gin.Context) {
	id := c.Param("meet_id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given meet_id is empty"})
		return
	}

	eventId, err := strconv.Atoi(c.Param("event_id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "event_id is not of type number"})
		return
	}

	event, err := service.GetEventByMeetingAndNumber(id, eventId)
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

func importEvent(c *gin.Context) {
	var request dto.ImportEventRequestDto
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event, r, err := service.ImportEvent(request.Event, request.StyleName, request.MeetingPartNumber)
	if err != nil {
		fmt.Printf(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if r {
		c.IndentedJSON(http.StatusCreated, event)
	} else {
		c.IndentedJSON(http.StatusOK, event)
	}
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
