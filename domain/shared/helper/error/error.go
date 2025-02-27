package error

import (
	"context"
	"database/sql"
	"encoding/json"
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

func New(statusCode int, msg string, err error) error {
	return fmt.Errorf("%d | %s | %w", statusCode, msg, err)
}

func TrimMessage(err error) (statusCode int, customError, originalError string) {
	errs := strings.Split(err.Error(), "|")
	statusCode, _ = strconv.Atoi(strings.TrimSpace(errs[0]))
	customError = strings.TrimSpace(errs[1])
	originalError = strings.TrimSpace(errs[2])
	return
}

func ResponseErrorWithContext(ctx context.Context, err error, slackNotif slack.SlackNotificationBug) error {
	var (
		statusCode    int
		customError   string
		originalError string
	)

	statusCode, customError, originalError = TrimMessage(err)
	body := ctx.Value(constant.FiberContext).(*fiber.Ctx).Body()
	method := ctx.Value(constant.FiberContext).(*fiber.Ctx).Method()
	var jsonBody map[string]interface{}
	if strings.ToLower(method) != "get" {
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			log.Printf("Error unmarshaling body request to JSON: %v", err)
			jsonBody = map[string]interface{}{"body": string(body)}
		}
	}

	if statusCode != 404 && slackNotif != nil {
		err = slackNotif.Send(fmt.Sprintf("ERROR CODE: %d, ERROR MESSAGE: %s, BODY REQUEST: %s", statusCode, originalError, jsonBody))
		if err != nil {
			originalError = originalError + ", ERROR SEND NOTIFICATION SLACK: " + err.Error()
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

func HandleError(err error) error {
	switch {
	case err == context.DeadlineExceeded:
		return New(http.StatusInternalServerError, "timeout", err)
	case err == sql.ErrNoRows:
		return New(http.StatusNotFound, "data not found", err)
	default:
		return New(http.StatusInternalServerError, "something when wrong", err)
	}
}
