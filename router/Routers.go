package router

import (
    "belajar-go/controllers"
    "belajar-go/middleware"

    "github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
    r.POST("/register", controllers.SignUp)
    r.POST("/login", controllers.Login)
    r.GET("/validate", middleware.RequireAuth, controllers.Validate)
}
