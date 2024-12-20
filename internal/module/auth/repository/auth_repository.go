package repository

import (
	"context"
	"errors"

	"github.com/evrintobing17/dating-app-go/internal/models"
	"github.com/evrintobing17/dating-app-go/internal/module/auth"
	"github.com/evrintobing17/dating-app-go/internal/repository"
	"github.com/lib/pq"
)

type authRepo struct {
	db *repository.Database
}

func NewAuthRepository(db *repository.Database) auth.AuthRepository {
	return &authRepo{db: db}
}

func (a *authRepo) Login(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, email, password, is_premium FROM users WHERE email=$1"
	err := a.db.Conn.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsPremium)

	return user, err
}

func (a *authRepo) SignUp(query string, user *models.User) error {
	_, err := a.db.Conn.Exec(query, user.Name, user.Email, user.Password, user.IsPremium)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return errors.New("user already exists")
			}
		}
	}
	return err
}

func (a *authRepo) UpdateUser(context context.Context, userID int) error {
	query := "UPDATE users SET is_premium = TRUE WHERE id = $1"
	_, err := a.db.Conn.Exec(query, userID)
	return err
}
