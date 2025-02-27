package helper

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/domain/shared/model"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Validate(request interface{}) error {
	validate := validator.New()
	// Validate the request struct
	if err := validate.Struct(request); err != nil {
		// Validation failed
		var errReq string

		// Extract individual validation errors
		for _, err := range err.(validator.ValidationErrors) {
			errReq = fmt.Sprintf("Field '%s' is required\n", camelCaseToSpaces(err.Field()))
			break
		}
		err = sharedError.New(http.StatusBadRequest, errReq, err)
		return err
	}
	return nil
}

func camelCaseToSpaces(s string) string {
	var result strings.Builder

	for i, c := range s {
		if i > 0 && unicode.IsUpper(c) {
			result.WriteRune(' ')
		}
		result.WriteRune(unicode.ToLower(c))
	}

	return result.String()
}

// Secret key (harus disimpan dengan aman, misalnya di env)
var secretKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken creates a JWT token with user_id and role claims
func GenerateToken(userID string, role string) (string, error) {
	// Set token claims
	claims := model.Claim{
		Role:   role,
		UserId: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	return token.SignedString(secretKey)
}

// ValidateToken parses and validates a JWT token
func ValidateToken(tokenString string) (model.Claim, error) {
	// Parse token with claims
	token, err := jwt.ParseWithClaims(tokenString, &model.Claim{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return model.Claim{}, fmt.Errorf("failed to parse token: %w", err)
	}

	// Extract claims
	if claims, ok := token.Claims.(*model.Claim); ok && token.Valid {
		// Check if the token has expired
		if claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
			return model.Claim{}, errors.New("token has expired")
		}

		return *claims, nil
	}

	return model.Claim{}, errors.New("invalid token")
}

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	// Generate hash with default cost (recommended: bcrypt.DefaultCost = 10)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	// Convert to string and return
	return string(hashedBytes), nil
}

// VerifyPassword compares a hashed password with a plain text password
func VerifyPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil // Jika tidak ada error, berarti password cocok
}
