package handler

import (
	"net/http"
	"zayyid-go/domain/shared/context"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedResponse "zayyid-go/domain/shared/response"
	"zayyid-go/domain/third_party/model"

	"github.com/gofiber/fiber/v2"
)

func (h ThirdPartyHandler) CallbackPaymentReceivingMidtrans(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req model.MidtransNotificationBodyReq
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	err = h.feature.MidtransNotificationFeature(ctx, req)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}

func (h ThirdPartyHandler) FrontendPaymentNotification(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req model.FrontendNotificationBodyReq
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	err = h.feature.FrontendPaymentNotificationFeature(ctx, req)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}
