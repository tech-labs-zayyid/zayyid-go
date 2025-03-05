package feature

import (
	"context"
	sharedHelperErr "zayyid-go/domain/shared/helper/error"
	sharedHelper "zayyid-go/domain/shared/helper/general"
	sharedHelperRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/domain/user/model"
)

func (f UserFeature) CreateAgentFeature(ctx context.Context, payload model.RegisterRequest, userId string) (resp model.UserRes, err error) {

	// if role not agent return error
	if payload.Role != "agent" {
		err = sharedHelperErr.New(403, "Role should be an agent", nil)
		return
	}

	// if agent was register generate referal code
	referalCode, errGenerateReferal := sharedHelper.GenerateRandomString(10)
	if errGenerateReferal != nil {
		err = errGenerateReferal
		return
	}

	payload.ReferalCode = referalCode

	// generate random string for password
	password, err := sharedHelper.GenerateRandomString(5)
	if err != nil {
		return
	}

	// encrypt password
	encryptedPassword, err := sharedHelper.HashPassword(password)
	if err != nil {
		return
	}

	// set up password
	payload.Password = encryptedPassword

	// generate agent id
	agentId := sharedHelperRepo.GenerateUuidAsIdTable()

	// call repo
	err = f.repo.RegisterRepository(ctx, payload, agentId.String())
	if err != nil {
		return
	}

	// get one user by userid
	user, err := f.repo.GetUserById(ctx, agentId.String())
	if err != nil {
		return
	}

	// Generate token
	token, err := sharedHelper.GenerateToken(userId, payload.Role)
	if err != nil {
		return
	}

	// generate refresh token
	refreshToken, err := sharedHelper.GenerateRefreshToken(user.Id, user.Role)
	if err != nil {
		return
	}

	// TODO: send password to agent email
	resp = model.UserRes{
		Id:             agentId.String(),
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
