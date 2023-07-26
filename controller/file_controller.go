package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/swimresults/meeting-service/model"
	"github.com/swimresults/meeting-service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func fileController() {
	router.GET("/file", getFiles)
	router.GET("/file/:id", getFile)
	router.GET("/file/meeting/list/:meeting", getFileListByMeeting)
	router.GET("/file/meeting/:meeting/name/:name", getFileByNameAndMeeting)
	router.DELETE("/file/:id", removeFile)
	router.POST("/file", addFile)
	router.POST("/file/increment", incrementDownloads)
	router.PUT("/file", updateFile)
}

func getFiles(c *gin.Context) {
	files, err := service.GetFiles()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, files)
}

func getFileListByMeeting(c *gin.Context) {
	meeting := c.Param("meeting")
	if meeting == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given meeting was empty"})
		return
	}
	files, err := service.GetFileListByMeeting(meeting)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, files)
}

func getFile(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	file, err := service.GetFileById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, file)
}

func getFileByNameAndMeeting(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given name was empty"})
		return
	}
	meeting := c.Param("meeting")
	if meeting == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given meeting was empty"})
		return
	}

	file, err := service.GetFileByNameAndMeeting(name, meeting)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, file)
}

func removeFile(c *gin.Context) {
	id, convErr := primitive.ObjectIDFromHex(c.Param("id"))
	if convErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "given id was not of type ObjectID"})
		return
	}

	err := service.RemoveFileById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, "")
}

func addFile(c *gin.Context) {
	var file model.StorageFile
	if err := c.BindJSON(&file); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.AddFile(file)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

func updateFile(c *gin.Context) {
	var file model.StorageFile
	if err := c.BindJSON(&file); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	r, err := service.UpdateFile(file)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

func incrementDownloads(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	path := string(data[:])

	f, err2 := service.IncrementDownloads(path)
	if err2 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err2.Error()})
		return
	}

	if f {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusNotFound)
	}

}
