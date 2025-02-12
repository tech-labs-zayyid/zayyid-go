package handler

import (
	sharedResponse "middleware-cms-api/domain/shared/response"
	"middleware-cms-api/domain/user_menu/feature"
	"middleware-cms-api/domain/user_menu/model"
	"middleware-cms-api/infrastructure/logger"

	"github.com/gofiber/fiber/v2"
)

type integrationMenuHandler struct {
	feature         *feature.UserMenuFeature
	isRequestLogged bool
}

func NewUserMenuHandler(feature *feature.UserMenuFeature, isRequestLogged bool) UserHandlerInterface {
	return &integrationMenuHandler{
		feature:         feature,
		isRequestLogged: isRequestLogged,
	}
}

func (h integrationMenuHandler) GetAppTypeHandler(c *fiber.Ctx) error {

	appType := []string{
		"API",
		"DB",
		"STORAGE",
		"MESSAGING",
		"SAP",
	}

	response := model.AppType{
		Data:       appType,
		StatusCode: 200,
	}

	// Return raw body json as is from the target server
	return c.Status(response.StatusCode).JSON(response)
}

func (h integrationMenuHandler) GetListMenuHandler(c *fiber.Ctx) error {
	var request model.Menu
	if err := c.QueryParser(&request); err != nil {
		return sharedResponse.BadRequestError(c, "Bad Request : "+err.Error(), "", nil)
	}

	if h.isRequestLogged {
		logger.LogInfoWithData(request, "REQUEST_LOG", "Incoming Request")
	}

	response, err := h.feature.GetListMenu()
	if err != nil {
		return err
	}

	// Return raw body json as is from the target server
	return c.Status(response.StatusCode).JSON(response)
}

func (h integrationMenuHandler) GetListHandler(c *fiber.Ctx) error {
	var request model.User
	if err := c.QueryParser(&request); err != nil {
		return sharedResponse.BadRequestError(c, "Bad Request : "+err.Error(), "", nil)
	}

	if h.isRequestLogged {
		logger.LogInfoWithData(request, "REQUEST_LOG", "Incoming Request")
	}

	response, err := h.feature.GetList(request)
	if err != nil {
		return err
	}

	// Return raw body json as is from the target server
	return c.Status(response.StatusCode).JSON(response)
}

func (h integrationMenuHandler) GetDataByIdHandler(c *fiber.Ctx) error {
	var request model.User
	id := c.Params("id")
	if id == "" {
		return sharedResponse.BadRequestError(c, "Bad Request : id is nil", "", nil)
	}
	request.Id = id

	if h.isRequestLogged {
		logger.LogInfoWithData(request, "REQUEST_LOG", "Incoming Request")
	}

	response, err := h.feature.GetDataById(request.Id)
	if err != nil {
		return err
	}

	// Return raw body json as is from the target server
	return c.Status(response.StatusCode).JSON(response)
}

func (h integrationMenuHandler) CreateDataHandler(c *fiber.Ctx) error {
	var request model.User
	if err := c.BodyParser(&request); err != nil {
		return sharedResponse.BadRequestError(c, "Bad Request : "+err.Error(), "", nil)
	}

	if h.isRequestLogged {
		logger.LogInfoWithData(request, "REQUEST_LOG", "Incoming Request")
	}

	// response, err := h.feature.UpsertData(c.Context(), request)
	response, err := h.feature.UpsertDataManual(c.Context(), request)
	if err != nil {
		return err
	}

	// Return raw body json as is from the target server
	return c.Status(response.StatusCode).JSON(response)
}

func (h integrationMenuHandler) UpdateDataHandler(c *fiber.Ctx) error {
	var request model.User
	if err := c.BodyParser(&request); err != nil {
		return sharedResponse.BadRequestError(c, "Bad Request : "+err.Error(), "", nil)
	}

	if h.isRequestLogged {
		logger.LogInfoWithData(request, "REQUEST_LOG", "Incoming Request")
	}

	response, err := h.feature.UpsertData(c.Context(), request)
	if err != nil {
		return err
	}

	// Return raw body json as is from the target server
	return c.Status(response.StatusCode).JSON(response)
}

func (h integrationMenuHandler) DeleteDataHandler(c *fiber.Ctx) error {
	var request model.User
	if err := c.BodyParser(&request); err != nil {
		return sharedResponse.BadRequestError(c, "Bad Request : "+err.Error(), "", nil)
	}

	if h.isRequestLogged {
		logger.LogInfoWithData(request, "REQUEST_LOG", "Incoming Request")
	}

	response, err := h.feature.DeleteData(c.Context(), request.Id)
	if err != nil {
		return err
	}

	// Return raw body json as is from the target server
	return c.Status(response.StatusCode).JSON(response)
}

func (h integrationMenuHandler) ActivatedHandler(c *fiber.Ctx) error {
	var request model.User
	if err := c.BodyParser(&request); err != nil {
		return sharedResponse.BadRequestError(c, "Bad Request : "+err.Error(), "", nil)
	}

	if h.isRequestLogged {
		logger.LogInfoWithData(request, "REQUEST_LOG", "Incoming Request")
	}

	response, err := h.feature.Activate(c.Context(), request)
	if err != nil {
		return err
	}

	// Return raw body json as is from the target server
	return c.Status(response.StatusCode).JSON(response)
}
