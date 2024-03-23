package router

import (
	"belajar-go/controllers"
	"belajar-go/middleware"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	// Register
	r.POST("/users/register", controllers.SignUp)
	r.GET("/register", func(c *gin.Context) {
		c.File("./views/Register.html")
	})

	// Login
	r.POST("/users/login", controllers.Login)
	r.GET("/login", func(c *gin.Context) {
		c.File("./views/Login.html")
	})

	// Middleware for all routes
	r.NoRoute(middleware.RequireAuth)

	m := r.Group("/")
	m.Use(middleware.RequireAuth)
	{
		// Homepage
		m.GET("/", func(c *gin.Context) {
			c.File("./views/Homepage.html")
		})

		// User Things
		/////////////////////////////////////////////////////
		// Get User Info
		m.GET("/users/info", controllers.GetUserInfo)

		// Edit Users
		m.GET("/users", func(c *gin.Context) {
			c.File("./views/Users.html")
		})
		m.PUT("/users/:userId", controllers.EditUsers)

		// Logout
		m.POST("/users/logout", controllers.Logout)

		// Delete User
		m.DELETE("/users/:userId", controllers.DeleteUser)
		/////////////////////////////////////////////////////

		// Photo Things
		/////////////////////////////////////////////////////
		// Add Photo
		m.POST("/photos", controllers.PostPhoto)

		// Get Photo
		m.GET("/photos", controllers.GetPhoto)

		// Get InfoPhoto
		m.GET("/infophoto", controllers.GetInfoPhoto)

		// Edit Photo
		m.PUT("/photos/:photoId", controllers.EditPhotos)

		// Delete Photo
		m.DELETE("/photos/:photoId", controllers.DeletePhoto)
		/////////////////////////////////////////////////////

	}
}
