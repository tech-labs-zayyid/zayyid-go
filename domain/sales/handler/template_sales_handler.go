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

// Add Data Template godoc
// @Summary      Add Data Template
// @Description  add data of Template
// @Tags         Data Template
// @Param        payload    body   request.AddTemplateReq  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/gallery [post]
func (h salesHandler) AddTemplateSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req request.AddTemplateReq
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	if err = h.feature.AddTemplateSales(ctx, req); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}

// Get List Tamplate godoc
// @Summary      Get List Template
// @Description  show list of Template
// @Tags         Data Template
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/temlate [get]
func (h salesHandler) GetListTemplateSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	resp, err := h.feature.GetListTemplateSales(ctx)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Get List Public Tamplate godoc
// @Summary      Get List Public Template
// @Description  show list of Public Template
// @Tags         Data Template
// @param        subdomain path string true "subdomain"
// @param        referral path string true "referral"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /public/temlate/{subdomain}/{referral} [get]
func (h salesHandler) GetListPublicTemplateSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	subdomain := c.Params("subdomain")
	if subdomain == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	referral := c.Params("*")
	resp, err := h.feature.GetListPublicTemplateSales(ctx, subdomain, referral)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Get Detail Template godoc
// @Summary      Get Detail Template
// @Description  show detail of Template
// @Tags         Data Template
// @Param        id path string false "id"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/template/{id} [get]
func (h salesHandler) GetTemplateSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	resp, err := h.feature.GetDetailTemplateSales(ctx, id)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Update Data Template godoc
// @Summary      Update Data Template
// @Description  Update Data of Template
// @Tags         Data Template
// @Param        id path string false "id"
// @Param        payload    body   request.UpdateTemplateReq  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/template/{id} [put]
func (h salesHandler) UpdateTemplateSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req request.UpdateTemplateReq
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	req.Id = c.Params("id")
	if req.Id == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	if err = h.feature.UpdateTemplateSales(ctx, req); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}
