package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/swimresults/meeting-service/model"
	"github.com/swimresults/meeting-service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func meetingController() {
	router.GET("/meeting", getMeetings)
	router.GET("/meeting/:id", getMeeting)
	router.GET("/meeting/meet/:meet_id", getMeetingByMeetId)
	router.GET("/meeting/between/:date_start/:date_end", getMeetingWithDateBetween)
	router.DELETE("/meeting/:id", removeMeeting)
	router.POST("/meeting", addMeeting)
	router.PUT("/meeting", updateMeeting)
}

func getMeetings(c *gin.Context) {
	meetings, err := service.GetMeetings()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, meetings)
}

func getMeeting(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	meeting, err := service.GetMeetingById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, meeting)
}

func getMeetingByMeetId(c *gin.Context) {
	id := c.Param("meet_id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given meet_id is empty"})
		return
	}

	meeting, err := service.GetMeetingByMeetId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, meeting)
}

func getMeetingWithDateBetween(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}

func removeMeeting(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	err := service.RemoveMeetingById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, "")
}

func addMeeting(c *gin.Context) {
	var meeting model.Meeting
	if err := c.BindJSON(&meeting); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.AddMeeting(meeting)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

func updateMeeting(c *gin.Context) {
	var meeting model.Meeting
	if err := c.BindJSON(&meeting); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.UpdateMeeting(meeting)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}
