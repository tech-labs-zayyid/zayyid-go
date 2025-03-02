package feature

import (
	"context"
	sharedHelper "zayyid-go/domain/shared/helper"
	sharedHelperRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/domain/user/model"
)

func (f UserFeature) RegisterFeature(ctx context.Context, payload model.RegisterRequest) (resp model.UserRes, err error) {

	// Encrypt password using hash
	encryptedPassword, err := sharedHelper.HashPassword(payload.Password)
	if err != nil {
		return
	}

	// override actual password
	payload.Password = encryptedPassword

	userId := sharedHelperRepo.GenerateUuidAsIdTable()

	// if agent was register generate referal code
	if payload.Role == "agent" {
		referalCode, errGenerateReferal := sharedHelper.GenerateRandomString(10)
		if errGenerateReferal != nil {
			err = errGenerateReferal
			return
		}

		payload.ReferalCode = referalCode
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

	// generate refresh token
	refreshToken, err := sharedHelper.GenerateRefreshToken(user.Id, user.Role)
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
			Token:        token,
			RefreshToken: refreshToken,
		},
	}

	return

}
