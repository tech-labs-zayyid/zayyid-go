package handler

import (
	"zayyid-go/domain/third_party/feature"

	"github.com/gofiber/fiber/v2"
)

type ThirdPartyHandler struct {
	feature *feature.ThirdPartyFeature
}

func NewThirdPartyHandler(feature *feature.ThirdPartyFeature) ThirdPartyHandler {
	return ThirdPartyHandler{
		feature: feature,
	}
}

type ThirdPartyHandlerInterface interface {
	CallbackPaymentReceivingMidtrans(c *fiber.Ctx) (err error)
}
