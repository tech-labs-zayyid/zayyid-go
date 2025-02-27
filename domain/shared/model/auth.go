package model

import "github.com/golang-jwt/jwt/v5"

// Claim struct untuk JWT
type Claim struct {
	Role   string `json:"role"`
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}
