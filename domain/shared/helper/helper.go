package helper

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"
	sharedError "zayyid-go/domain/shared/helper/error"

	"github.com/go-playground/validator/v10"
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
