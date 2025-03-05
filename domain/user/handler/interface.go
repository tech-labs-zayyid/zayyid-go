package handler

import (
	"zayyid-go/domain/user/feature"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	feature *feature.UserFeature
}

func NewUserHandler(feature *feature.UserFeature) UserHandler {
	return UserHandler{
		feature: feature,
	}
}

type IUserHandler interface {
	RegisterUserHandler(c *fiber.Ctx) (err error)
	AuthUserHandler(c *fiber.Ctx) (err error)
	RefreshTokenHandler(c *fiber.Ctx) (err error)
	UpdateHandler(c *fiber.Ctx) (err error)
	CreateAgentHandler(c *fiber.Ctx) (err error)
	GetAgentHandler(c *fiber.Ctx) (err error)
}
