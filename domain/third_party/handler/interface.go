package handler

import (
	"zayyid-go/domain/third_party/feature"
	"zayyid-go/infrastructure/service/slack"

	"github.com/gofiber/fiber/v2"
)

type ThirdPartyHandler struct {
	feature   *feature.ThirdPartyFeature
	slackConf slack.SlackNotificationBug
}

func NewThirdPartyHandler(feature *feature.ThirdPartyFeature, slackConf slack.SlackNotificationBug) ThirdPartyHandler {
	return ThirdPartyHandler{
		feature:   feature,
		slackConf: slackConf,
	}
}

type ThirdPartyHandlerInterface interface {
	CallbackPaymentReceivingMidtrans(c *fiber.Ctx) (err error)
}
