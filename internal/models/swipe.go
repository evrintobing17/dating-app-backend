package models

type Swipe struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProfileID int    `json:"profile_id"`
	Action    string `json:"action"`
	SwipedAt  string `json:"swiped_at"`
}
