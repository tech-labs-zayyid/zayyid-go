package middleware

import (
	"fmt"

	sharedConstant "middleware-cms-api/domain/shared/helper/constant"

	"github.com/golang-jwt/jwt/v4"
)

func CheckToken(tokenString string) (err error) {
	// Parse the token without verification
	token, err := jwt.Parse(tokenString, nil)
	if err != nil && err.Error() != sharedConstant.ERRJWTVALIDATION {
		return err
	}

	err = nil

	// Extract data from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid claim token")
	}

	// sso v3
	sub := claims["sub"] // is user_id
	if sub == 0 {
		return fmt.Errorf("invalid Authorization header format")
	}

	return
}
