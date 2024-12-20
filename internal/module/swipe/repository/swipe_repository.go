package repository

import (
	"context"
	"time"

	"github.com/evrintobing17/dating-app-go/internal/models"
	"github.com/evrintobing17/dating-app-go/internal/module/swipe"
	"github.com/evrintobing17/dating-app-go/internal/repository"
)

type swipeRepo struct {
	db *repository.Database
}

func NewSwipeRepository(db *repository.Database) swipe.SwipeRepository {
	return &swipeRepo{db: db}
}

func (a *swipeRepo) Login(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, email, password, is_premium FROM users WHERE email=$1"
	err := a.db.Conn.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsPremium)

	return user, err
}

// CreateSwipe implements swipe.SwipeRepository.
func (a *swipeRepo) CreateSwipe(userID int, profileID int, action string, time time.Time) error {
	query := "INSERT INTO swipes (user_id, profile_id, action, swiped_at) VALUES ($1, $2, $3, $4)"
	_, err := a.db.Conn.Exec(query, userID, profileID, action, time)
	return err
}
