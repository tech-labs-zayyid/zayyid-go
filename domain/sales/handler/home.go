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

// Get Data Home godoc
// @Summary      Get Data Home
// @Description  show list data of Home
// @Tags         Public
// @param        subdomain path string true "subdomain"
// @param        referral path string false "referral"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /public/home/{subdomain}/{referral} [get]
func (h salesHandler) GetDataHome(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	subdomain := c.Params("subdomain")
	if subdomain == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	referral := c.Params("*")
	resp, err := h.feature.HomeSalesData(ctx, subdomain, referral)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}
