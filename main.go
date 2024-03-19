package main

import (
	"belajar-go/database"
	"belajar-go/helpers"
	"belajar-go/router"

	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnvVariables()
	database.Connect()
	database.Migration()
}

func main() {
	r := gin.Default()
	router.Routers(r)
	r.Run()
}
