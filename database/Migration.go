package database

import (
	"belajar-go/models"
)

func Migration() {
	DB.AutoMigrate(&models.User{})
}
