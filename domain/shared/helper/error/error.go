package error

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	ctxShared "zayyid-go/domain/shared/context"
	"zayyid-go/domain/shared/helper/constant"
	"zayyid-go/infrastructure/logger"
	"zayyid-go/infrastructure/service/slack"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

func TrimMessage(err error) (statusCode int, customError string) {
	errs := strings.Split(err.Error(), ":")
	statusCode, _ = strconv.Atoi(strings.TrimSpace(errs[0]))
	customError = strings.TrimSpace(errs[1])

	return
}

func ResponseErrorWithContext(ctx context.Context, err error, slackNotif slack.SlackNotificationBug) error {
	var (
		statusCode    int
		customError   string
		originalError string
	)

	statusCode, customError = TrimMessage(err)
	body := ctx.Value(constant.FiberContext).(*fiber.Ctx).Body()
	method := ctx.Value(constant.FiberContext).(*fiber.Ctx).Method()
	var jsonBody map[string]interface{}

	if strings.ToLower(method) != "get" {
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			log.Printf("Error unmarshaling body request to JSON: %v", err)
			jsonBody = map[string]interface{}{"body": string(body)}
		}
	}

	if statusCode != http.StatusNotFound {
		if slackNotif != nil {
			if err := slackNotif.Send(fmt.Sprintf("ERROR CODE: %d, ERROR MESSAGE: %s, BODY REQUEST: %v", statusCode, originalError, jsonBody)); err != nil {
				originalError += ", ERROR SEND NOTIFICATION SLACK: " + err.Error()
			}
		}
	}

	logger.LogError(constant.RESPONSE, customError, originalError)
	response := Response{
		Status:  constant.ERROR,
		Message: customError,
		Data:    nil,
	}

	c := ctxShared.GetValueFiberFromContext(ctx)

	return c.Status(statusCode).JSON(response)
}

type CustomError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

func New(statusCode int, message string, err error) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

var (
	ErrInvalidEmailPassword = New(http.StatusUnauthorized, "invalid email or password", nil)
	ErrInvalidToken         = New(http.StatusBadRequest, "invalid token", nil)
	ErrMissingJWTSecret     = New(http.StatusInternalServerError, "missing JWT secret", nil)
	ErrFailedToParseToken   = New(http.StatusBadRequest, "failed to parse token", nil)
	ErrInvalidHeaderFormat  = New(http.StatusBadRequest, "invalid header format", nil)
)

func HandleError(err error) *CustomError {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return New(http.StatusInternalServerError, "timeout", err)
	case errors.Is(err, sql.ErrNoRows):
		return New(http.StatusNotFound, "data not found", err)
	case errors.Is(err, ErrInvalidEmailPassword):
		return ErrInvalidEmailPassword
	case errors.Is(err, ErrInvalidToken):
		return ErrInvalidToken
	case errors.Is(err, ErrMissingJWTSecret):
		return ErrMissingJWTSecret
	case errors.Is(err, ErrFailedToParseToken):
		return ErrFailedToParseToken
	case errors.Is(err, ErrInvalidHeaderFormat):
		return ErrInvalidHeaderFormat
	default:
		return New(http.StatusInternalServerError, "something went wrong", err)
	}
}
