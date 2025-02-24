package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/shared/context"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedResponse "zayyid-go/domain/shared/response"
)

// Add Data Social Media godoc
// @Summary      Add Data Social Media
// @Description  add data of Social Media
// @Tags         Data Social Media
// @Param        payload    body   request.AddSocialMediaReq  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/social-media [post]
func (h salesHandler) AddSocialMediaSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req request.AddSocialMediaReq
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	if len(req.DataSocialMedia) == 0 {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrRequestGallery, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	if err = h.feature.AddSocialMediaSales(ctx, req); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}

// Get List Social Media godoc
// @Summary      Get List Social Media
// @Description  show list of Social Media
// @Tags         Data Social Media
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/social-media [get]
func (h salesHandler) GetListSocialMediaSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	resp, err := h.feature.GetListSocialMediaSales(ctx)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Get List Public Social Media godoc
// @Summary      Get List Public Social Media
// @Description  show list of Public Social Media
// @Tags         Data Social Media
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @param        subdomain path string true "subdomain"
// @param        referral path string true "referral"
// @Router       /public/social-media/{subdomain}/{referral} [get]
func (h salesHandler) GetListPublicSocialMediaSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	subdomain := c.Params("subdomain")
	if subdomain == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	referral := c.Params("*")
	resp, err := h.feature.GetListSocialMediaPublicSales(ctx, subdomain, referral)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}
