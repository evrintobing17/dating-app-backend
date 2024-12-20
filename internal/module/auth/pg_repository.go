package auth

import (
	"context"

	"github.com/evrintobing17/dating-app-go/internal/models"
)

type AuthRepository interface {
	Login(ctx context.Context, email string) (*models.User, error)
	SignUp(query string, user *models.User) error
	UpdateUser(ctx context.Context, userID int) error
}
