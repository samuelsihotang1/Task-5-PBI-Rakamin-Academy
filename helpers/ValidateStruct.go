package helpers

import (
	"net/http"

	validate "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func MyValidateStruct(c *gin.Context, body interface{}) bool {
	// Validate request
	validateResult, err := validate.ValidateStruct(body)

	if err != nil && !validateResult {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to validate body",
		})
		return false
	} else {
		return true
	}
}
