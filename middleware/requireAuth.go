package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is middleware",
	})
	c.Next()
}