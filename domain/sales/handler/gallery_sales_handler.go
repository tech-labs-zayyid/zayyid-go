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

// Add Data Gallery godoc
// @Summary      Add Data Gallery
// @Description  add data of Gallery
// @Tags         Data Gallery
// @Param 		 Authorization header string true "Bearer token"
// @Param        payload    body   request.AddGalleryParam  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/gallery [post]
func (h salesHandler) AddGallerySales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req request.AddGalleryParam
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	if len(req.ImageUrl) == 0 {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrRequestGallery, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	if err = h.feature.AddGallerySales(ctx, req); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}

// Get List Gallery godoc
// @Summary      Get List Gallery
// @Description  show List of Gallery
// @Tags         Data Gallery
// @Param 		 Authorization header string true "Bearer token"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/gallery [get]
func (h salesHandler) GetGallerySales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	resp, err := h.feature.GetDataListGallery(ctx)
	if err != nil {
		err = sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
		return
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Get Detail Gallery godoc
// @Summary      Get Detail Gallery
// @Description  show detail of Gallery
// @Tags         Public
// @param        subdomain path string true "subdomain"
// @param        referral path string true "referral"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /public/gallery/{subdomain} [get]
func (h salesHandler) GetGallerySalesPublic(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	subdomain := c.Params("subdomain")
	if subdomain == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	resp, err := h.feature.GetDataListGalleryPublic(ctx, subdomain)
	if err != nil {
		err = sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
		return
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Get Detail Gallery godoc
// @Summary      Get Detail Gallery
// @Description  show detail of Gallery
// @Tags         Data Gallery
// @Param 		 Authorization header string true "Bearer token"
// @Param        id path string false "id"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/gallery/{id} [get]
func (h salesHandler) GetDataGallerySales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	resp, err := h.feature.GetDataGallerySales(ctx, id)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Update Data Gallery godoc
// @Summary      Update Data Gallery
// @Description  Update Data of Gallery
// @Tags         Data Gallery
// @Param 		 Authorization header string true "Bearer token"
// @Param        id path string false "id"
// @Param        payload    body   request.UpdateGalleryParam  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /sales/gallery/{id} [put]
func (h salesHandler) UpdateGallerySales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req request.UpdateGalleryParam
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	req.Id = c.Params("id")
	if req.Id == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	if err = h.feature.UpdateGallery(ctx, req); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}
