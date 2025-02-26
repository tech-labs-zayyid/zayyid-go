package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"zayyid-go/domain/shared/context"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedResponse "zayyid-go/domain/shared/response"
)

func (h salesHandler) GetDataHome(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	subdomain := c.Params("subdomain")
	if subdomain == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	referral := c.Params("*")
	resp, err := h.feature.HomeSalesData(ctx, subdomain, referral)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}
