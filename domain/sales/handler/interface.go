package handler

import (
	"github.com/gofiber/fiber/v2"
)

type SalesHandlerInterface interface {
	AddGallerySales(c *fiber.Ctx) (err error)
	GetGallerySales(c *fiber.Ctx) (err error)
	GetGallerySalesPublic(c *fiber.Ctx) (err error)
	AddBannerSales(c *fiber.Ctx) (err error)
	GetBannerSales(c *fiber.Ctx) (err error)
	AddBannerPublicSales(c *fiber.Ctx) (err error)
}
