package handler

import (
	"zayyid-go/domain/user/model"

	"github.com/gofiber/fiber/v2"
)

// UpdateUserHandler godoc
// @Summary Update a user
// @Description Update user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.UserReq true "User Data"
// @Success 200 {object} model.UserRes
// @Failure 400 {object} fiber.Error
// @Failure 404 {object} fiber.Error
// @Failure 500 {object} fiber.Error
// @Router /users/{id} [put]
func (h UserHandler) UpdateUserHandler(c *fiber.Ctx) (err error) {

	var userReq model.UserReq

	if err := c.BodyParser(&userReq); err != nil {
		return user, fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	// Assuming you have a method to update the user in your handler
	user, err = h.UpdateUser(id, userReq)
	if err != nil {
		if err == model.ErrUserNotFound {
			return user, fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	return nil

}
