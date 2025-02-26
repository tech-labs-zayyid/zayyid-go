package feature

import (
	"context"
	sharedHelper "zayyid-go/domain/shared/helper"
	"zayyid-go/domain/user/model"

	"github.com/google/uuid"
)

func (f UserFeature) RegisterFeature(ctx context.Context, payload model.RegisterRequest) (resp model.UserRes, err error) {

	// Encrypt password using hash
	encryptedPassword, err := sharedHelper.HashPassword(payload.Password)
	if err != nil {
		return
	}

	// override actual password
	payload.Password = encryptedPassword

	userId, err := uuid.NewV7()
	if err != nil {
		return
	}

	// call repo
	err = f.repo.RegisterRepository(ctx, payload, userId.String())
	if err != nil {
		return
	}

	// get one user by userid
	user, err := f.repo.GetUserById(ctx, userId.String())
	if err != nil {
		return
	}

	// Generate token
	token, err := sharedHelper.GenerateToken(userId.String(), payload.Role)
	if err != nil {
		return
	}

	resp = model.UserRes{
		Id:             userId.String(),
		Name:           user.Name,
		UserName:       user.UserName,
		Email:          user.Email,
		Role:           user.Role,
		WhatsAppNumber: user.WhatsAppNumber,
		CreatedAt:      user.CreatedAt,
		CreatedBy:      user.CreatedBy,
		TokenData: model.TokenRes{
			Token: token,
		},
	}

	return

}
