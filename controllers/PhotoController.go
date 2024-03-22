package controllers

import (
	"belajar-go/app"
	"belajar-go/database"
	"belajar-go/helpers"
	"belajar-go/models"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PostPhoto(c *gin.Context) {

	// Get request body
	body := app.Photo{}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Validate request
	if !helpers.MyValidateStruct(c, body) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to validate body",
		})
		return
	}

	// File Photo
	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	user, _ := c.Get("user")
	userID := strconv.Itoa(int(user.(models.User).ID))
	pathfile := "./images/" + userID + "_" + time.Now().Format("2006-01-02-15-04-05") + filepath.Ext(file.Filename)

	if err := c.SaveUploadedFile(file, pathfile); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	// Insert To Database
	photo := models.Photo{Title: body.Title, Caption: body.Caption, PhotoUrl: pathfile, UserID: user.(models.User).ID, User: user.(models.User)}
	result := database.DB.Create(&photo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create Photo",
		})

		return
	}

	c.String(http.StatusOK, "Photo %s uploaded successfully with Title=%s and Caption=%s.", file.Filename, body.Title, body.Caption)
}
