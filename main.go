package main

import (
	"task-5-pbi-btpns-SamuelChristyAngieSihotang/database"
	"task-5-pbi-btpns-SamuelChristyAngieSihotang/helpers"
	"task-5-pbi-btpns-SamuelChristyAngieSihotang/router"

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
