package helpers

import (
	validate "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func MyValidateStruct(c *gin.Context, body interface{}) bool {
	// Validate request
	validateResult, err := validate.ValidateStruct(body)

	if err != nil && !validateResult {
		return false
	} else {
		return true
	}
}
