package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/swimresults/meeting-service/model"
	"github.com/swimresults/meeting-service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func meetingSeriesController() {
	router.GET("/meeting_series", getMeetingSeries)
	router.GET("/meeting_series/:id", getMeetingSeriesById)
	router.DELETE("/meeting_series/:id", removeMeetingSeries)
	router.POST("/meeting_series", addMeetingSeries)
	router.PUT("/meeting_series", updateMeetingSeries)
}

func getMeetingSeries(c *gin.Context) {
	meetings, err := service.GetMeetingSeries()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, meetings)
}

func getMeetingSeriesById(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	meeting, err := service.GetMeetingSeriesById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, meeting)
}

func removeMeetingSeries(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	err := service.RemoveMeetingSeriesById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, "")
}

func addMeetingSeries(c *gin.Context) {
	var meeting model.MeetingSeries
	if err := c.BindJSON(&meeting); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.AddMeetingSeries(meeting)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

func updateMeetingSeries(c *gin.Context) {
	var meeting model.MeetingSeries
	if err := c.BindJSON(&meeting); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.UpdateMeetingSeries(meeting)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}
