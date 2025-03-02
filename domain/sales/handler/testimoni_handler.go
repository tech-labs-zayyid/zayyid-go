package handler

import (
	"errors"
	"net/http"
	modelRequest "zayyid-go/domain/sales/model/request"
	sharedHelper "zayyid-go/domain/shared/helper"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/domain/shared/response"

	"github.com/gofiber/fiber/v2"
)

// Create Data Testimoni godoc
// @Summary      Create Data Testimoni
// @Description  create data of Testimoni
// @Tags         Data Testimoni
// @Param        payload    body   modelRequest.Testimoni  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /testimony [post]
func (h salesHandler) AddTestimoniHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq modelRequest.Testimoni
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	errResp := sharedHelper.Validate(bodyReq)
	if errResp != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	err = h.feature.AddTestimoniFeature(ctx, bodyReq)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return response.ResponseOK(c, sharedConstant.SUCCESS, "")
}

// Update Data Testimoni godoc
// @Summary      Update Data Testimoni
// @Description  Update Data of Testimoni
// @Tags         Data Testimoni
// @Param        payload    body   modelRequest.Testimoni  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /testimony [put]
func (h salesHandler) UpdateTestimoniHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq modelRequest.Testimoni
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	errResp := sharedHelper.Validate(bodyReq)
	if errResp != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	err = h.feature.UpdateTestimoniFeature(ctx, bodyReq)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return response.ResponseOK(c, sharedConstant.SUCCESS, "")
}

// Get Detail Customer godoc
// @Summary      Get Detail Customer
// @Description  show detail of Customer
// @Tags         Data Testimoni
// @Param        id   query      string  false  "id"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /testimony [get]
func (h salesHandler) GetTestimoniHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	filter := new(modelRequest.Testimoni)

	if err := c.QueryParser(filter); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	data, err := h.feature.GetTestimoniFeature(ctx, *filter)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return response.ResponseOK(c, sharedConstant.SUCCESS, data)
}

// Get List Customer godoc
// @Summary      Get List Customer
// @Description  show list of Customer
// @Tags         Data Testimoni
// @Param        id   query      string  false  "id"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /testimony/list [get]
func (h salesHandler) GetListTestimoniHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	param := new(modelRequest.Testimoni)

	if err := c.QueryParser(param); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	filter := new(modelRequest.TestimoniSearch)

	if err := c.QueryParser(filter); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	data, pagination, err := h.feature.GetListTestimoniFeature(ctx, *param, *filter)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return response.ResponseOkWithPagination(c, sharedConstant.SUCCESS, data, pagination)
}

// Get Public List Customer godoc
// @Summary      Get Public public List Customer
// @Description  show list of Customer
// @Tags         Data Testimoni
// @Param        subdomain   path      string  true  "sub domain"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /testimony/list/{subdomain}/{referral} [get]
func (h salesHandler) GetPublicListTestimoniHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	referral := c.Params("*")
	subDomain := c.Params("subdomain")
	if subDomain == "" {
		err := sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	filter := new(modelRequest.TestimoniSearch)

	if err := c.QueryParser(filter); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	data, pagination, err := h.feature.GetPublicListTestimoniFeature(ctx, subDomain, referral, *filter)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return response.ResponseOkWithPagination(c, sharedConstant.SUCCESS, data, pagination)
}
