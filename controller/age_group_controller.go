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

func ageGroupController() {
	router.GET("/age_group", getAgeGroups)
	router.GET("/age_group/:id", getAgeGroup)
	router.GET("/age_group/meet/:meet_id", getAgeGroupsByMeeting)
	router.GET("/age_group/meet/:meet_id/event/:event_id", getAgeGroupByMeetingAndEvent)

	router.POST("/age_group", addAgeGroup)
	router.POST("/age_group/import", importAgeGroup)

	router.DELETE("/age_group/:id", removeAgeGroup)

	router.PUT("/age_group", updateAgeGroup)
}

func getAgeGroups(c *gin.Context) {
	ageGroups, err := service.GetAgeGroups()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, ageGroups)
}

func getAgeGroup(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	ageGroup, err := service.GetAgeGroupById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, ageGroup)
}

func getAgeGroupsByMeeting(c *gin.Context) {
	id := c.Param("meet_id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given meet_id is empty"})
		return
	}
	ageGroups, err := service.GetAgeGroupsByMeeting(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, ageGroups)
}

func getAgeGroupByMeetingAndEvent(c *gin.Context) {
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

	ageGroup, err := service.GetAgeGroupsByMeetingAndEvent(id, eventId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, ageGroup)
}

func removeAgeGroup(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	err := service.RemoveAgeGroupById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, "")
}

func importAgeGroup(c *gin.Context) {
	var request dto.ImportAgeGroupRequestDto
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ageGroup, r, err := service.ImportAgeGroup(request.AgeGroup)
	if err != nil {
		fmt.Printf(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if r {
		c.IndentedJSON(http.StatusCreated, ageGroup)
	} else {
		c.IndentedJSON(http.StatusOK, ageGroup)
	}
}

func addAgeGroup(c *gin.Context) {
	var ageGroup model.AgeGroup
	if err := c.BindJSON(&ageGroup); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.AddAgeGroup(ageGroup)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

func updateAgeGroup(c *gin.Context) {
	var ageGroup model.AgeGroup
	if err := c.BindJSON(&ageGroup); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.UpdateAgeGroup(ageGroup)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}
