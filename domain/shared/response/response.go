package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Url     string      `json:"url"`
	Data    interface{} `json:"data,omitempty"`
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
