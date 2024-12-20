package models

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserID    int  `json:"user_id"`
	IsPremium bool `json:"is_premium"`
	jwt.StandardClaims
}