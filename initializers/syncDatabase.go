package initializers

import "belajar-go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}