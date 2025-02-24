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

// Add Data Banner godoc
// @Summary      Add Data Banner
// @Description  add data of Banner
// @Tags         Data Banner
// @Param        payload    body   request.BannerReq  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/banner [post]
func (h salesHandler) AddBannerSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req request.BannerReq
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	if len(req.DataBanner) == 0 {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrRequestBanner, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	if err = h.feature.AddBannerSales(ctx, req); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}

// Get List Banner godoc
// @Summary      Get List Banner
// @Description  show list of Banner
// @Tags         Data Banner
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/banner [get]
func (h salesHandler) GetListBannerSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	resp, err := h.feature.GetListBannerSales(ctx)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Get Public Banner godoc
// @Summary      Get Public Banner
// @Description  show public of Banner
// @Tags         Data Banner
// @param        subdomain path string true "subdomain"
// @param        referral path string true "referral"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /public/banner/{subdomain}/{refferal} [get]
func (h salesHandler) GetBannerPublicSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	subdomain := c.Params("subdomain")
	if subdomain == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	referral := c.Params("*")
	resp, err := h.feature.GetListBannerPublic(ctx, subdomain, referral)
	if err != nil {
		err = sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
		return
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Get Detail Banner godoc
// @Summary      Get Detail Banner
// @Description  show detail of Banner
// @Tags         Data Banner
// @Param        id path string false "id"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/banner/{id} [get]
func (h salesHandler) GetBannerSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	resp, err := h.feature.GetBannerSales(ctx, id)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Update Data Banner godoc
// @Summary      Update Data Banner
// @Description  Update Data of Banner
// @Tags         Data Banner
// @param        id path string true "id"
// @Param        payload    body   request.BannerUpdateReq  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/banner/{id} [put]
func (h salesHandler) UpdateBannerSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req request.BannerUpdateReq
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	req.Id = c.Params("id")
	if req.Id == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	if err = h.feature.UpdateBanner(ctx, req); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}
