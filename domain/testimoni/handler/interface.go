package handler

import "github.com/gofiber/fiber/v2"

type UserHandlerInterface interface {
	Ping(c *fiber.Ctx) error
}
