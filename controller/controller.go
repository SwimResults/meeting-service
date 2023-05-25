package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sr-meeting/meeting-service/service"
)

var router = gin.Default()

func Run() {

	port := os.Getenv("SR_MEETING_PORT")

	if port == "" {
		fmt.Println("no application port given! Please set SR_MEETING_PORT.")
		return
	}

	meetingController()
	meetingSeriesController()
	locationController()
	styleController()
	eventController()

	router.GET("/actuator", actuator)

	err := router.Run(":" + port)
	if err != nil {
		fmt.Println("Unable to start application on port " + port)
		return
	}
}

func actuator(c *gin.Context) {

	state := "OPERATIONAL"

	if !service.PingDatabase() {
		state = "DATABASE_DISCONNECTED"
	}
	c.String(http.StatusOK, state)
}
