package handler

import "github.com/gofiber/fiber/v2"

type TestimoniInterface interface {
	Ping(c *fiber.Ctx) error

	AddTestimoniHandler(c *fiber.Ctx) error
	UpdateTestimoniHandler(c *fiber.Ctx) error
	GetTestimoniHandler(c *fiber.Ctx) error
	GetListTestimoniHandler(c *fiber.Ctx) error
}
