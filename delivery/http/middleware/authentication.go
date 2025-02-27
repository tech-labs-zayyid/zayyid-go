package middleware

import (
	"errors"
	"fmt"
	"os"

	sharedModel "zayyid-go/domain/shared/model"

	"github.com/golang-jwt/jwt/v5"
)

func MiddlewareAuth(tokenString string) error {
	if tokenString == "" {
		return fmt.Errorf("missing token")
	}

	// Ambil secret key dari environment
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return fmt.Errorf("missing JWT_SECRET environment variable")
	}

	// Parse token dengan kunci
	token, err := jwt.ParseWithClaims(tokenString, &sharedModel.Claim{}, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode signing sesuai
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return fmt.Errorf("invalid token signature")
		}
		return fmt.Errorf("failed to parse token: %w", err)
	}

	// Ambil data claims
	claims, ok := token.Claims.(*sharedModel.Claim)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid claim token")
	}

	// Validasi UserId untuk SSO V3
	if claims.UserId == "" {
		return fmt.Errorf("invalid Authorization header format")
	}

	return nil
}
