package feature

import (
	"context"
	"net/http"
	sharedHelper "zayyid-go/domain/shared/helper"
	sharedHelperErr "zayyid-go/domain/shared/helper/error"
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
		err = sharedHelperErr.New(http.StatusBadRequest, "Wrong email or password", err)
		return
	}

	// generate token when success login
	token, err := sharedHelper.GenerateToken(user.Id, user.Role)
	if err != nil {
		return
	}

	// generate refresh token
	refreshToken, err := sharedHelper.GenerateRefreshToken(user.Id, user.Role)
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
			Token:        token,
			RefreshToken: refreshToken,
		},
	}

	return

}

func (f UserFeature) RefreshTokenFeature(ctx context.Context, refreshToken string) (resp model.TokenRes, err error) {

	// validate the refresh token
	claims, err := sharedHelper.ValidateToken(refreshToken)
	if err != nil {
		err = sharedHelperErr.New(http.StatusBadRequest, "Bad request", err)
		return
	}

	// get user by id from token claims
	user, err := f.repo.GetUserById(ctx, claims.UserId)
	if err != nil {
		return
	}

	// generate new token
	newToken, err := sharedHelper.GenerateToken(user.Id, user.Role)
	if err != nil {
		return
	}

	resp = model.TokenRes{
		Token:        newToken,
		RefreshToken: refreshToken,
	}

	return

}
