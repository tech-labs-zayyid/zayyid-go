package handler

import (
	"net/http"
	sharedHelper "zayyid-go/domain/shared/helper"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/domain/shared/response"
	"zayyid-go/domain/testimoni/feature"
	"zayyid-go/domain/testimoni/model"

	"github.com/gofiber/fiber/v2"
)

type testimoniHandler struct {
	feature         *feature.TestimoniFeature
	isRequestLogged bool
}

func NewTestimoniHandler(feature *feature.TestimoniFeature, isRequestLogged bool) TestimoniInterface {
	return &testimoniHandler{
		feature:         feature,
		isRequestLogged: isRequestLogged,
	}
}

func (h testimoniHandler) Ping(c *fiber.Ctx) error {
	response := http.StatusText(http.StatusOK)

	h.feature.Ping(c.Context())

	return c.Status(http.StatusOK).JSON(response)
}

// Create Data Testimoni godoc
// @Summary      Create Data Testimoni
// @Description  create data of Testimoni
// @Tags         Data Testimoni
// @Param        payload    body   model.Testimoni  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /testimoni [post]
func (h testimoniHandler) AddTestimoniHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq model.Testimoni
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	errResp := sharedHelper.Validate(bodyReq)
	if errResp != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	err = h.feature.UpsertTestimoniFeature(ctx, bodyReq)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return response.ResponseOK(c, sharedConstant.SUCCESS, "")
}

// Update Data Testimoni godoc
// @Summary      Update Data Testimoni
// @Description  Update Data of Testimoni
// @Tags         Data Testimoni
// @Param        payload    body   model.Testimoni  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /testimoni [put]
func (h testimoniHandler) UpdateTestimoniHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq model.Testimoni
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	errResp := sharedHelper.Validate(bodyReq)
	if errResp != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	bodyReq.IsUpdate = 1
	err = h.feature.UpsertTestimoniFeature(ctx, bodyReq)
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
// @Router       /testimoni [get]
func (h testimoniHandler) GetTestimoniHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	filter := new(model.Testimoni)

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
// @Router       /testimoni/list [get]
func (h testimoniHandler) GetListTestimoniHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	filter := new(model.Testimoni)

	if err := c.QueryParser(filter); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	data, pagination, err := h.feature.GetListTestimoniFeature(ctx, *filter)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return response.ResponseOkWithPagination(c, sharedConstant.SUCCESS, data, pagination)
}
