package middleware

import (
	"fmt"
	"net/http"
	"os"
	"task-5-pbi-btpns-SamuelChristyAngieSihotang/database"
	"task-5-pbi-btpns-SamuelChristyAngieSihotang/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie using the cookie name 'token'
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
		// c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/Validate it
	// Parse takes the token string and a function for looking up the key. The latter is especially
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
		// log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Redirect(http.StatusFound, "/login")
			return
			// c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with that id
		var user models.User
		resultSearch := database.DB.First(&user, claims["sub"])

		if resultSearch.RowsAffected == 0 {
			c.Redirect(http.StatusFound, "/login")
			return
		}

		if resultSearch.Error != nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}

		// Attach to request
		c.Set("user", user)

		// Continue
		c.Next()
	} else {
		// Memaksa untuk mengakses route /login
		c.Redirect(http.StatusFound, "/login")
		return
		// c.AbortWithStatus(http.StatusUnauthorized)
	}
}
