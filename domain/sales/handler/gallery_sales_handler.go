package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"zayyid-go/domain/sales/feature"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/shared/context"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedResponse "zayyid-go/domain/shared/response"
)

type salesHandler struct {
	feature         *feature.SalesFeature
	isRequestLogged bool
}

func NewSalesHandler(feature *feature.SalesFeature, isRequestLogged bool) SalesHandlerInterface {
	return &salesHandler{
		feature:         feature,
		isRequestLogged: isRequestLogged,
	}
}

func (h salesHandler) AddGallerySales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req request.AddGalleryParam
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	if err = h.feature.AddGallerySales(ctx, req); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}

func (h salesHandler) GetGallerySales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	resp, err := h.feature.GetDataListGallery(ctx)
	if err != nil {
		err = sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
		return
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

func (h salesHandler) GetGallerySalesPublic(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	subdomain := c.Params("subdomain")
	if subdomain == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	resp, err := h.feature.GetDataListGalleryPublic(ctx, subdomain)
	if err != nil {
		err = sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
		return
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}
