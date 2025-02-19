package response

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strings"
	"zayyid-go/domain/shared/helper/constant"
	"zayyid-go/domain/shared/model"
	"zayyid-go/infrastructure/logger"
)

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Url     string      `json:"url"`
	Data    interface{} `json:"data,omitempty"`
}

type NewResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

func NetworkConnectionError(c *fiber.Ctx, message string, url string) error {
	return c.Status(fiber.StatusServiceUnavailable).JSON(Response{
		Status:  "Service Unavailable",
		Code:    fiber.StatusServiceUnavailable,
		Message: message,
		Url:     url,
	})
}

func InternalServerError(c *fiber.Ctx, message string, url string, data interface{}) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Status:  "Internal Server Error",
		Code:    fiber.StatusInternalServerError,
		Message: message,
		Url:     url,
		Data:    data,
	})
}

func NotFoundError(c *fiber.Ctx, message string, url string, data interface{}) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Status:  "Not Found",
		Code:    fiber.StatusNotFound,
		Message: message,
		Url:     url,
		Data:    data,
	})
}

func BadRequestError(c *fiber.Ctx, message string, url string, data interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status:  "Bad Request",
		Code:    fiber.StatusBadRequest,
		Message: message,
		Url:     url,
		Data:    data,
	})
}

func NotAuthorizedError(c *fiber.Ctx, message string, url string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Status:  "Unauthorized",
		Code:    fiber.StatusUnauthorized,
		Message: message,
		Url:     url,
	})
}

func RequestTimeoutError(c *fiber.Ctx, message string, url string) error {
	return c.Status(fiber.StatusRequestTimeout).JSON(Response{
		Status:  "Request Timeout",
		Code:    fiber.StatusRequestTimeout,
		Message: message,
		Url:     url,
	})
}

func ResponseCreatedOK(c *fiber.Ctx, msg, customMessage string, data interface{}) error {
	body := c.Body()
	method := c.Method()
	var jsonBody map[string]interface{}
	if strings.ToLower(method) != "get" {
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			log.Printf("Error unmarshaling body request to JSON: %v", err)
			jsonBody = map[string]interface{}{"body": string(body)}
		}
	}

	logger.LogInfoWithData(data, constant.RESPONSE, msg)
	response := Response{
		Status:  constant.SUCCESS,
		Message: msg,
		Data:    data,
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func ResponseOkWithPagination(c *fiber.Ctx, msg string, data interface{}, pagination interface{}) error {
	logger.LogInfoWithData(data, constant.RESPONSE, msg)
	response := model.ResponseWithPagination{
		Status:     constant.SUCCESS,
		Message:    msg,
		Data:       data,
		Pagination: pagination,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func ResponseOK(c *fiber.Ctx, msg string, data interface{}) error {
	body := c.Body()
	method := c.Method()
	var jsonBody map[string]interface{}
	if strings.ToLower(method) != "get" {
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			log.Printf("Error unmarshaling body request to JSON: %v", err)
			jsonBody = map[string]interface{}{"body": string(body)}
		}
	}

	logger.LogInfoWithData(data, constant.RESPONSE, msg)
	response := NewResponse{
		Status:  constant.SUCCESS,
		Message: msg,
		Data:    data,
	}

	return c.Status(http.StatusOK).JSON(response)
}
