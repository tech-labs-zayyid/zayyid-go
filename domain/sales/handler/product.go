package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/shared/context"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedPaginate "zayyid-go/domain/shared/helper/pagination"
	sharedModel "zayyid-go/domain/shared/model"
	sharedResponse "zayyid-go/domain/shared/response"
)

// Add Data Banner godoc
// @Summary      Add Data Banner
// @Description  add data of Banner
// @Tags         Data Product
// @Param 		 Authorization header string true "Bearer token"
// @Param        payload    body   request.AddProductReq  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Router       /sales/product [post]
func (h salesHandler) AddProductSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req request.AddProductReq
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	if err = h.feature.AddProductSales(ctx, req); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}

// Get List Product godoc
// @Summary      Get List Product
// @Description  show List of Product
// @Tags         Data Product
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Param 		 Authorization header string true "Bearer token"
// @Router       /sales/product [get]
func (h salesHandler) GetListProductSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	search := strings.TrimSpace(c.Query(sharedConstant.SEARCH))
	isActive := strings.TrimSpace(c.Query(sharedConstant.IS_ACTIVE))
	subCategoryProduct := strings.TrimSpace(c.Query(sharedConstant.SUB_CATEGORY_PRODUCT))
	bestProduct := strings.TrimSpace(c.Query(sharedConstant.BEST_PRODUCT))
	statusProduct := strings.TrimSpace(c.Query(sharedConstant.STATUS_PRODUCT))
	sortBy := strings.TrimSpace(c.Query(sharedConstant.SORT_BY))
	sortOrder := strings.TrimSpace(c.Query(sharedConstant.SORT_ORDER))
	page := sharedPaginate.GetPageOrDefault(c.Query(sharedConstant.PAGE), 1)
	limit := sharedPaginate.GetLimitOrDefault(c.Query(sharedConstant.LIMIT), 20)

	queryRequest := sharedModel.QueryRequest{
		Search:             search,
		SubCategoryProduct: subCategoryProduct,
		BestProduct:        bestProduct,
		StatusProduct:      statusProduct,
		IsActive:           isActive,
		Page:               page,
		Limit:              limit,
		SortBy:             sortBy,
		SortOrder:          sortOrder,
	}

	resp, pagination, err := h.feature.ListProductSales(ctx, queryRequest)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOkWithPagination(c, sharedConstant.SUCCESS, resp, pagination)
}

// Get Detail Product godoc
// @Summary      Get Detail Product
// @Description  show Detail of Product
// @Tags         Data Product
// @param        id path string true "id"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Param 		 Authorization header string true "Bearer token"
// @Router       /sales/product/{id} [get]
func (h salesHandler) GetDetailProductSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	resp, err := h.feature.GetDetailSalesProduct(ctx, id)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), resp)
}

// Update Product godoc
// @Summary      Update Product
// @Description  update of Product
// @Tags         Data Product
// @param        id path string true "id"
// @Param        payload    body   request.UpdateProductSales  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Param 		 Authorization header string true "Bearer token"
// @Router       /sales/product/{id} [put]
func (h salesHandler) UpdateProductSales(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var req request.UpdateProductSales
	if err = c.BodyParser(&req); err != nil {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, err)
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	id := c.Params("id")
	if id == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	req.ProductId = id
	err = h.feature.UpdateProductSales(ctx, req)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, http.StatusText(http.StatusOK), "")
}

// Get List Product Public godoc
// @Summary      Get List Product Public
// @Description  show List of Product Public
// @Tags         Public
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /public/product/{domain} [get]
func (h salesHandler) GetListProductSalesPublic(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	var param request.ProductListPublic
	err = c.QueryParser(&param)
	if err != nil {
		return
	}

	subdomain := c.Params("subdomain")
	if subdomain == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	resp, pagination, err := h.feature.GetListProductSalesPublic(ctx, param, subdomain)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOkWithPagination(c, sharedConstant.SUCCESS, resp, pagination)
}

// Get Detail Product Public godoc
// @Summary      Get Detail Product Public
// @Description  show Detail of Product Public
// @Tags         Public
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /public/product/detai/{domain}/{slug} [get]
func (h salesHandler) GetDetailProductSalesPublic(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	subdomain := c.Params("subdomain")
	if subdomain == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	slug := c.Params("slug")
	if subdomain == "" {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrInvalidRequest, errors.New(sharedConstant.ErrInvalidRequest))
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	resp, err := h.feature.DetailProductSalesPublic(ctx, subdomain, slug)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.slackConf)
	}

	return sharedResponse.ResponseOK(c, sharedConstant.SUCCESS, resp)
}
