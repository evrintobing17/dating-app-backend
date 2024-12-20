package middleware

import (
	"net/http"
	"strings"

	"github.com/evrintobing17/dating-app-go/internal/module/auth"
	"github.com/evrintobing17/dating-app-go/internal/utils"
	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	authService auth.AuthRepository
}

func NewAuthMiddleware(authService auth.AuthRepository) AuthMiddleware {
	return &authMiddleware{authService: authService}
}

func (auth *authMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}
		jwtTokenSplit := strings.Split(token, "Bearer ")
		if jwtTokenSplit[1] == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		jwtToken := jwtTokenSplit[1]

		claims, err := utils.ParseToken(jwtToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Attach claims to context
		c.Set("userID", claims.UserID)
		c.Set("isPremium", claims.IsPremium)

		c.Next()
	}
}
