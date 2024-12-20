package utils

import (
	"github.com/evrintobing17/dating-app-go/internal/models"
	"github.com/gin-gonic/gin"
)

// JSONResponse is a helper to format JSON responses consistently.
func JSONResponse(c *gin.Context, code int, message string, data interface{}, errors interface{}) {
	c.JSON(code, models.BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
		Errors:  errors,
	})
}
