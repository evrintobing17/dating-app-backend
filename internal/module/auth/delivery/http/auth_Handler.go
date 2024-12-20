package http

import (
	"net/http"

	"github.com/evrintobing17/dating-app-go/internal/models"
	"github.com/evrintobing17/dating-app-go/internal/module/auth"
	"github.com/evrintobing17/dating-app-go/internal/utils"
	"github.com/evrintobing17/dating-app-go/pkg/validation"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase auth.AuthUsecase
}

func NewAuthHandler(r *gin.Engine, authUsecase auth.AuthUsecase) {
	handler := &AuthHandler{
		authUsecase: authUsecase,
	}
	authorized := r.Group("/auth")
	{
		authorized.POST("/login", handler.login)
		authorized.POST("/signup", handler.signup)
	}
}

func (h *AuthHandler) signup(c *gin.Context) {
	var req models.SignUpRequest
	if err := validation.ValidateRequest(c, &req); err != nil {
		return
	}

	user := &models.User{Name: req.Name, Email: req.Email, Password: req.Password}
	err := h.authUsecase.SignUp(c.Request.Context(), user)
	if err != nil {
		if err.Error() == "user already exists" {
			utils.JSONResponse(c, http.StatusBadRequest, "Bad Request", nil, err.Error())
			return
		}
		utils.JSONResponse(c, http.StatusInternalServerError, "Internal server error", nil, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "User created successfully", nil, nil)
}

func (h *AuthHandler) login(c *gin.Context) {
	var req models.LoginRequest

	if err := validation.ValidateRequest(c, &req); err != nil {
		return
	}

	token, err := h.authUsecase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Bad Request", nil, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Login successful", map[string]string{"token": token}, nil)
}
