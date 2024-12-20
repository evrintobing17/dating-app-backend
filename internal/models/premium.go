package models

type PremiumPurchase struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	PurchasedAt string `json:"purchased_at"`
}
