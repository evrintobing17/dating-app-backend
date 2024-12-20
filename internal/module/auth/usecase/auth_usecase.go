package usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/evrintobing17/dating-app-go/internal/models"
	"github.com/evrintobing17/dating-app-go/internal/module/auth"
	"github.com/evrintobing17/dating-app-go/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	repo auth.AuthRepository
}

func NewAuthUsecase(repo auth.AuthRepository) auth.AuthUsecase {
	return &AuthUseCase{repo: repo}
}

func (s *AuthUseCase) SignUp(ctx context.Context, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	query := "INSERT INTO users (name, email, password, is_premium) VALUES ($1, $2, $3, $4)"
	err = s.repo.SignUp(query, user)
	return err
}

func (s *AuthUseCase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.Login(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user Not Found")
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	return utils.GenerateToken(user.ID, user.IsPremium)
}
