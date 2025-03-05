package handler

import (
	"zayyid-go/domain/shared/context"
	sharedContext "zayyid-go/domain/shared/context"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedHelper "zayyid-go/domain/shared/helper/general"
	sharedResponse "zayyid-go/domain/shared/response"
	"zayyid-go/domain/user/model"

	"github.com/gofiber/fiber/v2"
)

// RegisterUserHandler godoc
// @Summary Create agent
// @Description Register a new user with the provided details
// @Tags Agent
// @Accept json
// @Produce json
// @Param payload body model.RegisterRequest true "Register Request"
// @Success 200 {object} sharedResponse.Response{data=model.UserRes}
// @Failure 400 {object} sharedResponse.Response
// @Failure 500 {object} sharedResponse.Response
// @Router /agent/create [post]
func (h UserHandler) CreateAgentHandler(c *fiber.Ctx) (err error) {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	// Define user model
	payload := model.RegisterRequest{}

	// Parse payload
	if err = c.BodyParser(&payload); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	// Validate payload
	err = sharedHelper.Validate(payload)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	// get user id from local fiber
	ctxValue := sharedContext.GetValueContext(ctx)
	userId := ctxValue.UserId

	// Call register feature
	resp, err := h.feature.CreateAgentFeature(ctx, payload, userID)
	if err != nil {
		// Handle for any error
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, "Create agent success!", resp)

}

// GetAgentHandler godoc
// @Summary Get agent list
// @Description Retrieve a list of agents based on the provided query parameters
// @Tags Agent
// @Accept json
// @Produce json
// @Param query query model.QueryAgentList true "Query Agent List"
// @Success 200 {object} sharedResponse.Response{data=[]model.UserRes}
// @Failure 400 {object} sharedResponse.Response
// @Failure 500 {object} sharedResponse.Response
// @Router /agent/list [get]
func (h UserHandler) GetAgentHandler(c *fiber.Ctx) (err error) {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	// Define user model
	query := model.QueryAgentList{}

	// Parse query
	if err = c.QueryParser(&query); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	// Validate query
	err = sharedHelper.Validate(query)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	// get user id from local fiber
	ctxValue := sharedContext.GetValueContext(ctx)
	userId := ctxValue.UserId

	// Call register feature
	resp, err := h.feature.GetAgentFeature(ctx, query, userId)
	if err != nil {
		// Handle for any error
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, "Get agent list success!", resp)
}