package handler

import "github.com/gofiber/fiber/v2"

type UserHandlerInterface interface {
	GetAppTypeHandler(c *fiber.Ctx) error
	GetListMenuHandler(c *fiber.Ctx) error
	GetListHandler(c *fiber.Ctx) error
	GetDataByIdHandler(c *fiber.Ctx) error
	CreateDataHandler(c *fiber.Ctx) error
	UpdateDataHandler(c *fiber.Ctx) error
	DeleteDataHandler(c *fiber.Ctx) error
	ActivatedHandler(c *fiber.Ctx) error
}
