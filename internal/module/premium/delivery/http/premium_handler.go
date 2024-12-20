package http

import (
	"net/http"

	"github.com/evrintobing17/dating-app-go/internal/middleware"
	"github.com/evrintobing17/dating-app-go/internal/module/premium"
	"github.com/evrintobing17/dating-app-go/internal/utils"
	"github.com/gin-gonic/gin"
)

type PremiumHandler struct {
	premiumUsecase premium.PremiumUsecase
	authMiddleware middleware.AuthMiddleware
}

func NewPremiumHandler(r *gin.Engine, premiumUsecase premium.PremiumUsecase, authMiddleware middleware.AuthMiddleware) {
	handler := &PremiumHandler{
		premiumUsecase: premiumUsecase,
		authMiddleware: authMiddleware,
	}
	authorized := r.Group("/premium", handler.authMiddleware.AuthMiddleware())
	{
		authorized.POST("/upgrade", handler.upgrade)
	}
}

func (h *PremiumHandler) upgrade(c *gin.Context) {
	userID := c.GetInt("userID")

	err := h.premiumUsecase.UpgradeToPremium(c.Request.Context(), userID)
	if err != nil {
		if err.Error() == "user already premium" {
			utils.JSONResponse(c, http.StatusBadRequest, "Bad Request", nil, err.Error())
			return
		}
		utils.JSONResponse(c, http.StatusInternalServerError, "Internal server error", nil, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "User created successfully", nil, nil)
}
