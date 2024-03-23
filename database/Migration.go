package database

import (
	"task-5-pbi-btpns-SamuelChristyAngieSihotang/models"
)

func Migration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Photo{})
}
