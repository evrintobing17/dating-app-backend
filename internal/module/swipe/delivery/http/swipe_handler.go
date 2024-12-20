package http

import (
	"net/http"

	"github.com/evrintobing17/dating-app-go/internal/middleware"
	"github.com/evrintobing17/dating-app-go/internal/models"
	"github.com/evrintobing17/dating-app-go/internal/module/swipe"
	"github.com/gin-gonic/gin"
)

type SwipeHandler struct {
	swipeUsecase   swipe.SwipeUsecase
	authMiddleware middleware.AuthMiddleware
}

func NewSwipeHandler(r *gin.Engine, swipeUsecase swipe.SwipeUsecase, authMiddleware middleware.AuthMiddleware) {
	handler := &SwipeHandler{
		swipeUsecase:   swipeUsecase,
		authMiddleware: authMiddleware,
	}
	authorized := r.Group("/", handler.authMiddleware.AuthMiddleware())
	{
		authorized.POST("swipe", handler.swipe)
	}
}

func (h *SwipeHandler) swipe(c *gin.Context) {
	var req models.SwipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt("userID")
	isPremium := c.GetBool("isPremium")

	if err := h.swipeUsecase.Swipe(c.Request.Context(), userID, req.ProfileID, req.Action, isPremium); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Swipe successful"})

}
