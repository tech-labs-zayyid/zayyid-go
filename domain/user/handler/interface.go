package handler

import (
	"zayyid-go/domain/user/feature"
	"zayyid-go/infrastructure/service/slack"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	feature   *feature.UserFeature
	SlackConf slack.SlackNotificationBug
}

func NewUserHandler(feature *feature.UserFeature) UserHandler {
	return UserHandler{
		feature: feature,
	}
}

type IUserHandler interface {
	RegisterUserHandler(c *fiber.Ctx) (err error)
	AuthUserHandler(c *fiber.Ctx) (err error)
}
