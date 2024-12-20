package swipe

import (
	"context"
)

type SwipeUsecase interface {
	Swipe(ctx context.Context, userID, profileID int, action string, isPremium bool) error
}
