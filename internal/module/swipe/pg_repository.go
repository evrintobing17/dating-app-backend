package swipe

import (
	"time"
)

type SwipeRepository interface {
	CreateSwipe(userID int, profileID int, action string, time time.Time) error
}
