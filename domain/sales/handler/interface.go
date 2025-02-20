package handler

import (
	"github.com/gofiber/fiber/v2"
)

type SalesHandlerInterface interface {
	AddGallerySales(c *fiber.Ctx) (err error)
	GetGallerySales(c *fiber.Ctx) (err error)
}
