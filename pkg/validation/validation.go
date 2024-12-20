package validation

import (
	"net/http"

	"github.com/evrintobing17/dating-app-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ValidateRequest validates the request body and returns an error response if invalid.
func ValidateRequest(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(req); err != nil {
		var validationErrors []map[string]string
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range errs {
				validationErrors = append(validationErrors, map[string]string{
					"field":   fieldErr.Field(),
					"message": fieldErr.Error(),
				})
			}
			utils.JSONResponse(c, http.StatusBadRequest, "Validation failed", nil, validationErrors)
			return err
		}
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid request body", nil, err.Error())
		return err
	}
	return nil
}
