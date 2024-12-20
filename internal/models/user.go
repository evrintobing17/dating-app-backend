
// internal/models/user.go
package models

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsPremium bool   `json:"is_premium"`
}

// internal/models/swipe.go

// internal/models/premium_purchase.go
