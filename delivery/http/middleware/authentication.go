package middleware

import (
	"context"
	"errors"
	"os"
	"strings"
	"zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedModel "zayyid-go/domain/shared/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// MiddlewareAuth untuk memvalidasi JWT sebelum request masuk ke handler
func Auth(c *fiber.Ctx) error {
	ctx := context.WithValue(context.Background(), constant.FiberContext, c)

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return sharedError.ResponseErrorWithContext(ctx, sharedError.HandleError(sharedError.ErrInvalidToken), nil)
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return sharedError.ResponseErrorWithContext(ctx, sharedError.ErrInvalidHeaderFormat, nil)
	}
	tokenString := tokenParts[1]

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return sharedError.ResponseErrorWithContext(ctx, sharedError.ErrMissingJWTSecret, nil)
	}

	token, err := jwt.ParseWithClaims(tokenString, &sharedModel.Claim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, sharedError.HandleError(sharedError.ErrFailedToParseToken)
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return sharedError.ResponseErrorWithContext(ctx, sharedError.ErrInvalidToken, nil)
		}
		return sharedError.ResponseErrorWithContext(ctx, sharedError.ErrFailedToParseToken, nil)
	}

	claims, ok := token.Claims.(*sharedModel.Claim)
	if !ok || !token.Valid {
		return sharedError.ResponseErrorWithContext(ctx, sharedError.ErrInvalidToken, nil)
	}

	if claims.UserId == "" {
		return sharedError.ResponseErrorWithContext(ctx, sharedError.ErrInvalidHeaderFormat, nil)
	}

	c.Locals("user_id", claims.UserId)
	c.Locals("role", claims.Role)

	return c.Next()

}
