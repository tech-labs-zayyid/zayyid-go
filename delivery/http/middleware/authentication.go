package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/domain/shared/helper/general"
	sharedModel "zayyid-go/domain/shared/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// MiddlewareAuth untuk memvalidasi JWT sebelum request masuk ke handler
func Auth(c *fiber.Ctx) error {

	ctx := context.WithValue(context.Background(), constant.FiberContext, c)

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return sharedError.ResponseErrorWithContext(ctx, sharedError.New(http.StatusBadRequest, "Invalid token", errors.New("authorization header is missing")), nil)
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return sharedError.ResponseErrorWithContext(ctx, sharedError.New(http.StatusBadRequest, "Invalid token", errors.New("invalid authorization format")), nil)
	}
	tokenString := tokenParts[1]

	token, err := jwt.ParseWithClaims(tokenString, &sharedModel.Claim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, sharedError.New(http.StatusBadRequest, "Failed to parse token", errors.New("unexpected signing method"))
		}
		return general.SecretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return sharedError.ResponseErrorWithContext(ctx, sharedError.New(http.StatusUnauthorized, "Invalid token", errors.New("signature verification failed")), nil)
		}
		return sharedError.ResponseErrorWithContext(ctx, sharedError.New(http.StatusBadRequest, "Failed to parse token", err), nil)
	}

	claims, ok := token.Claims.(*sharedModel.Claim)
	if !ok || !token.Valid {
		return sharedError.ResponseErrorWithContext(ctx, sharedError.New(http.StatusUnauthorized, "Invalid token", errors.New("token is either invalid or expired")), nil)
	}

	if claims.UserId == "" {
		return sharedError.ResponseErrorWithContext(ctx, sharedError.New(http.StatusBadRequest, "Invalid header format", errors.New("user ID is missing in token claims")), nil)
	}

	c.Locals("user_id", claims.UserId)
	c.Locals("role", claims.Role)

	return c.Next()

}
