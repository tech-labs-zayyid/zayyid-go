package handler

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"zayyid-go/domain/master/feature"
	"zayyid-go/domain/shared/context"
	"zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedPaginate "zayyid-go/domain/shared/helper/pagination"
	sharedModel "zayyid-go/domain/shared/model"
	"zayyid-go/domain/shared/response"
)

type masterHandler struct {
	feature         *feature.MasterFeature
	isRequestLogged bool
}

func NewMasterHandler(feature *feature.MasterFeature, isRequestLogged bool) MasterHandlerInterface {
	return &masterHandler{
		feature:         feature,
		isRequestLogged: isRequestLogged,
	}
}

func (h masterHandler) PingMaster(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()

	h.feature.Ping(ctx)
	return response.ResponseOK(c, constant.SUCCESS, "")
}

func (h masterHandler) MasterListProvince(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	search := strings.TrimSpace(c.Query(constant.SEARCH))
	isActive := strings.TrimSpace(c.Query(constant.STATUS))
	sortBy := strings.TrimSpace(c.Query(constant.SORT_BY))
	sortOrder := strings.TrimSpace(c.Query(constant.SORT_ORDER))
	page := sharedPaginate.GetPageOrDefault(c.Query(constant.PAGE), 1)
	limit := sharedPaginate.GetLimitOrDefault(c.Query(constant.LIMIT), 20)

	queryRequest := sharedModel.QueryRequest{
		Search:    search,
		Status:    isActive,
		Page:      page,
		Limit:     limit,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}

	resp, pagination, err := h.feature.MasterProvince(ctx, queryRequest)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return response.ResponseOkWithPagination(c, constant.SUCCESS, resp, pagination)
}

func (h masterHandler) MasterListCity(c *fiber.Ctx) (err error) {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	search := strings.TrimSpace(c.Query(constant.SEARCH))
	isActive := strings.TrimSpace(c.Query(constant.STATUS))
	provinceId := strings.TrimSpace(c.Query(constant.PROVINCE_ID))
	sortBy := strings.TrimSpace(c.Query(constant.SORT_BY))
	sortOrder := strings.TrimSpace(c.Query(constant.SORT_ORDER))
	page := sharedPaginate.GetPageOrDefault(c.Query(constant.PAGE), 1)
	limit := sharedPaginate.GetLimitOrDefault(c.Query(constant.LIMIT), 20)

	queryRequest := sharedModel.QueryRequest{
		Search:     search,
		ProvinceId: provinceId,
		Status:     isActive,
		Page:       page,
		Limit:      limit,
		SortBy:     sortBy,
		SortOrder:  sortOrder,
	}

	resp, pagination, err := h.feature.MasterCity(ctx, queryRequest)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return response.ResponseOkWithPagination(c, constant.SUCCESS, resp, pagination)
}
