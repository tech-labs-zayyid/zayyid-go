package handler

import (
	"github.com/gofiber/fiber/v2"
)

type SalesHandlerInterface interface {
	AddGallerySales(c *fiber.Ctx) (err error)
	GetGallerySales(c *fiber.Ctx) (err error)
	GetGallerySalesPublic(c *fiber.Ctx) (err error)
	AddTestimoniHandler(c *fiber.Ctx) error
	UpdateTestimoniHandler(c *fiber.Ctx) error
	GetTestimoniHandler(c *fiber.Ctx) error
	GetListTestimoniHandler(c *fiber.Ctx) error
	AddBannerSales(c *fiber.Ctx) (err error)
	GetBannerSales(c *fiber.Ctx) (err error)
	AddBannerPublicSales(c *fiber.Ctx) (err error)
}
