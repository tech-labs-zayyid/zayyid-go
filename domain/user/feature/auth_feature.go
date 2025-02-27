package feature

import (
	"context"
	"errors"
	sharedHelper "zayyid-go/domain/shared/helper"
	"zayyid-go/domain/user/model"
)

func (f UserFeature) AuthUserFeature(ctx context.Context, payload model.AuthUserRequest) (resp model.UserRes, err error) {

	// get user by email
	user, err := f.repo.GetByQueryRepository(ctx, model.QueryUser{
		Email: payload.Email,
	})
	if err != nil {
		return
	}

	// compare for the password
	if !sharedHelper.VerifyPassword(user.Password, payload.Password) {
		err = errors.New("invalid password")
		return
	}

	// generate token when success login
	token, err := sharedHelper.GenerateToken(user.Id, user.Role)
	if err != nil {
		return
	}

	resp = model.UserRes{
		Id:             user.Id,
		UserName:       user.UserName,
		Name:           user.Name,
		WhatsAppNumber: user.WhatsAppNumber,
		Email:          user.Email,
		Role:           user.Role,
		CreatedAt:      user.CreatedAt,
		CreatedBy:      user.CreatedBy,
		TokenData: model.TokenRes{
			Token: token,
		},
	}

	return

}
