package handler

import (
	"errors"
	"zayyid-go/domain/shared/context"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedHelper "zayyid-go/domain/shared/helper/general"
	sharedResponse "zayyid-go/domain/shared/response"
	"zayyid-go/domain/user/model"

	"github.com/gofiber/fiber/v2"
)

// UpdateUserHandler godoc
// @Summary Update a user
// @Description Update data sales
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param payload body model.UpdateUser true "User Data"
// @Success 200 {object} sharedResponse.Response
// @Failure 400 {object} sharedResponse.Response
// @Failure 404 {object} sharedResponse.Response
// @Failure 500 {object} sharedResponse.Response
// @Router /user/update [put]
func (h UserHandler) UpdateHandler(c *fiber.Ctx) (err error) {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	// Define user model
	payload := model.UpdateUser{}

	// Parse payload
	if err = c.BodyParser(&payload); err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	// Validate payload
	err = sharedHelper.Validate(payload)
	if err != nil {
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	// Get user_id from local fiber
	userId, ok := c.Locals("user_id").(string)
	if !ok {
		return sharedError.ResponseErrorWithContext(ctx, errors.New("invalid user_id type"), h.feature.SlackConf)
	}

	// Call register feature
	err = h.feature.UpdateFeature(ctx, payload, userId)
	if err != nil {
		// Handle for any error
		return sharedError.ResponseErrorWithContext(ctx, err, h.feature.SlackConf)
	}

	return sharedResponse.ResponseOK(c, "Update user success!", nil)

}
