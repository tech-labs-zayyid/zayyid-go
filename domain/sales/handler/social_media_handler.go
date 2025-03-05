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
// @Param 		 Authorization header string true "Bearer token"
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
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	if len(req.DataSocialMedia) == 0 {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrRequestGallery, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	if err = h.feature.AddSocialMediaSales(ctx, req); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}

// Get List Social Media godoc
// @Summary      Get List Social Media
// @Description  show list of Social Media
// @Tags         Data Social Media
// @Param 		 Authorization header string true "Bearer token"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/social-media [get]
func (h salesHandler) GetListSocialMediaSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	resp, err := h.feature.GetListSocialMediaSales(ctx)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Get List Public Social Media godoc
// @Summary      Get List Public Social Media
// @Description  show list of Public Social Media
// @Tags         Public
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @param        subdomain path string true "subdomain"
// @param        referral path string true "referral"
// @Router       /public/social-media/{subdomain} [get]
func (h salesHandler) GetListPublicSocialMediaSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	subdomain := c.Params("subdomain")
	if subdomain == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	resp, err := h.feature.GetListSocialMediaPublicSales(ctx, subdomain)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Get Detail Social Media godoc
// @Summary      Get Detail Social Media
// @Description  show detail of Public Social Media
// @Tags         Data Social Media
// @Param 		 Authorization header string true "Bearer token"
// @Param        id path string false "id"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/social-media/{id} [get]
func (h salesHandler) GetSocialMediaSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	resp, err := h.feature.GetDetailSocialMediaSales(ctx, id)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Update Social Media godoc
// @Summary      Update Social Media
// @Description  Update Social Media
// @Tags         Data Social Media
// @Param 		 Authorization header string true "Bearer token"
// @Param        id path string false "id"
// @Param        payload    body   request.UpdateSocialMediaSales  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/social-media/{id} [put]
func (h salesHandler) UpdateSocialMediaSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req request.UpdateSocialMediaSales
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	req.Id = c.Params("id")
	if req.Id == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	if err = h.feature.UpdateSocialMediaSales(ctx, req); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}
