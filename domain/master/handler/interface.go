package handler

import "github.com/gofiber/fiber/v2"

type MasterHandlerInterface interface {
	PingMaster(c *fiber.Ctx) (err error)
	MasterListProvince(c *fiber.Ctx) (err error)
	MasterListCity(c *fiber.Ctx) (err error)
}
