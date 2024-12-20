package repository

import (
	"context"
	"time"

	"github.com/evrintobing17/dating-app-go/internal/module/premium"
	"github.com/evrintobing17/dating-app-go/internal/repository"
)

type premiumRepo struct {
	db *repository.Database
}

func NewPremiumRepository(db *repository.Database) premium.PremiumRepository {
	return &premiumRepo{db: db}
}

// Update implements premium.PremiumRepository.
func (a *premiumRepo) Update(ctx context.Context, userID int) error {
	currDate := time.Now()
	query := "INSERT INTO premium_purchases (user_id, purchased_at) VALUES ($1, $2)"
	_, err := a.db.Conn.Exec(query, userID, currDate)
	return err
}

func (a *premiumRepo) GetPremiumByID(ctx context.Context, userID int) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM premium_purchases WHERE user_id=$1"
	err := a.db.Conn.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
