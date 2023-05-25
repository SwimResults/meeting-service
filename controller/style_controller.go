package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sr-meeting/meeting-service/model"
	"sr-meeting/meeting-service/service"
)

func styleController() {
	router.GET("/style", getStyles)
	router.GET("/style/:id", getStyle)
	router.GET("/style/name/:name", getStyleByName)
	router.DELETE("/style/:id", removeStyle)
	router.POST("/style", addStyle)
	router.PUT("/style", updateStyle)
}

func getStyles(c *gin.Context) {
	styles, err := service.GetStyles()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, styles)
}

func getStyle(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	style, err := service.GetStyleById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, style)
}

func getStyleByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given name was empty"})
		return
	}

	style, err := service.GetStyleByName(name)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, style)
}

func removeStyle(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	err := service.RemoveStyleById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, "")
}

func addStyle(c *gin.Context) {
	var style model.Style
	if err := c.BindJSON(&style); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.AddStyle(style)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

func updateStyle(c *gin.Context) {
	var style model.Style
	if err := c.BindJSON(&style); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.UpdateStyle(style)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}
