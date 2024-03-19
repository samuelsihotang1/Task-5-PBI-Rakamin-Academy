package initializers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error

	DB, err = gorm.Open(mysql.Open(os.Getenv("DB")), &gorm.Config{})

	if err != nil {
		panic(" Could not connect mysql DB ")
	}
}
