package handler

import (
	"zayyid-go/domain/shared/context"
	sharedHelper "zayyid-go/domain/shared/helper"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedResponse "zayyid-go/domain/shared/response"
	"zayyid-go/domain/user/model"

	"github.com/gofiber/fiber/v2"
)

// RegisterUserHandler godoc
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags User
// @Accept json
// @Produce json
// @Param payload body model.RegisterRequest true "Register Request"
// @Success 200 {object} sharedResponse.Response{data=model.UserRes}
// @Failure 400 {object} sharedResponse.Response
// @Failure 500 {object} sharedResponse.Response
// @Router /user/register [post]
func (h UserHandler) RegisterUserHandler(c *fiber.Ctx) (err error) {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	// Define user model
	payload := model.RegisterRequest{}

	// Parse payload
	if err = c.BodyParser(&payload); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.SlackConf)
	}

	// Validate payload
	err = sharedHelper.Validate(payload)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.SlackConf)
	}

	// Call register feature
	resp, err := h.feature.RegisterFeature(ctx, payload)
	if err != nil {
		// Handle for any error
		return sharedError.ResponseErrorWithContext(ctx, err, h.SlackConf)
	}

	return sharedResponse.ResponseOK(c, "Register user success!", resp)

}
