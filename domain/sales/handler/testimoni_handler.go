package handler

import (
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
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	errResp := sharedHelper.Validate(bodyReq)
	if errResp != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	err = h.feature.AddTestimoniFeature(ctx, bodyReq)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
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
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	errResp := sharedHelper.Validate(bodyReq)
	if errResp != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	err = h.feature.UpdateTestimoniFeature(ctx, bodyReq)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
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
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	data, err := h.feature.GetTestimoniFeature(ctx, *filter)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
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
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	filter := new(modelRequest.TestimoniSearch)

	if err := c.QueryParser(filter); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	data, pagination, err := h.feature.GetListTestimoniFeature(ctx, *param, *filter)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return response.ResponseOkWithPagination(c, sharedConstant.SUCCESS, data, pagination)
}
