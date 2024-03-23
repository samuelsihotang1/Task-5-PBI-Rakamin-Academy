package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"task-5-pbi-btpns-SamuelChristyAngieSihotang/app"
	"task-5-pbi-btpns-SamuelChristyAngieSihotang/database"
	"task-5-pbi-btpns-SamuelChristyAngieSihotang/helpers"
	"task-5-pbi-btpns-SamuelChristyAngieSihotang/models"
	"time"

	"github.com/gin-gonic/gin"
)

func PostPhoto(c *gin.Context) {

	userPhoto, _ := c.Get("user")
	currentUser := userPhoto.(models.User)

	database.DB.Preload("Photos").First(&currentUser, currentUser.ID)

	if len(currentUser.Photos) >= 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "You're already Post Photo",
		})
		return
	}

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
		c.String(http.StatusBadRequest, "Failed to get photo file")
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user information",
		})
		return
	}
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

func GetPhoto(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user information",
		})
		return
	}
	currentUser := user.(models.User)

	database.DB.Preload("Photos").First(&currentUser, currentUser.ID)

	if len(currentUser.Photos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No photo found for the current user",
		})
		return
	}

	firstPhoto := currentUser.Photos[0]
	filePath := filepath.Join("./images/", filepath.Base(firstPhoto.PhotoUrl))

	c.File(filePath)
}

func GetInfoPhoto(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user information",
		})
		return
	}
	currentUser := user.(models.User)

	database.DB.Preload("Photos").First(&currentUser, currentUser.ID)

	if len(currentUser.Photos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No photo found for the current user",
		})
		return
	}

	firstPhoto := currentUser.Photos[0]

	c.JSON(http.StatusOK, gin.H{
		"photo": firstPhoto,
	})
}

func DeletePhoto(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user information",
		})
		return
	}
	currentUser := user.(models.User)

	photoID := c.Param("photoId")

	var photoToDelete models.Photo
	if err := database.DB.First(&photoToDelete, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found",
		})
		return
	}

	if photoToDelete.UserID != currentUser.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You are not authorized to delete this photo",
		})
		return
	}

	if err := os.Remove(photoToDelete.PhotoUrl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete photo from filesystem",
		})
		return
	}

	if err := database.DB.Delete(&photoToDelete).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete photo from database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo deleted successfully",
	})
}

func EditPhotos(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user information",
		})
		return
	}
	currentUser := user.(models.User)

	photoID := c.Param("photoId")

	var photoToUpdate models.Photo
	if err := database.DB.First(&photoToUpdate, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found",
		})
		return
	}

	if photoToUpdate.UserID != currentUser.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You are not authorized to edit this photo",
		})
		return
	}

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

	// Update fields if they are provided in the request
	if body.Title != "" {
		photoToUpdate.Title = body.Title
	}
	if body.Caption != "" {
		photoToUpdate.Caption = body.Caption
	}

	// Check if a new file is uploaded
	file, err := c.FormFile("file")
	if err == nil {
		// Remove old file
		if err := os.Remove(photoToUpdate.PhotoUrl); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to delete old photo file",
			})
			return
		}

		// Save new file
		userID := strconv.Itoa(int(currentUser.ID))
		newFilePath := "./images/" + userID + "_" + time.Now().Format("2006-01-02-15-04-05") + filepath.Ext(file.Filename)
		if err := c.SaveUploadedFile(file, newFilePath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to upload new photo file",
			})
			return
		}
		photoToUpdate.PhotoUrl = newFilePath
	}

	// Update photo in database
	if err := database.DB.Save(&photoToUpdate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update photo in database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo updated successfully",
	})
}
