package feature

import (
	"context"
	"errors"
	"net/http"
	sharedHelperErr "zayyid-go/domain/shared/helper/error"
	sharedHelper "zayyid-go/domain/shared/helper/general"
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
		if payload.SalesId == "" {
			err = sharedHelperErr.New(http.StatusBadRequest, "sales id cannot be null", errors.New("bad request"))
			return
		}

		referalCode, errGenerateReferal := sharedHelper.GenerateRandomString(10)
		if errGenerateReferal != nil {
			err = sharedHelperErr.New(http.StatusBadRequest, "error generate referal", errGenerateReferal)
			return
		}

		payload.ReferalCode = referalCode

		// use transaction for registration
		trx := f.repo.OpenTransaction(ctx)

		// register with transaction
		err = f.repo.RegisterRepositoryTransaction(ctx, payload, userId.String(), trx)
		if err != nil {
			err = sharedHelperErr.HandleError(err)
			return
		}

		// mapping sales agent
		err = f.repo.MappingSalesAgent(ctx, payload.SalesId, userId.String(), payload.Email, trx)
		if err != nil {
			// rollback transaction
			trx.Rollback()
			err = sharedHelperErr.HandleError(err)
			return
		}

		// commit transaction
		trx.Commit()
	} else { // register as sales
		// call repo
		err = f.repo.RegisterRepository(ctx, payload, userId.String())
		if err != nil {
			return
		}
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
		TokenData: &model.TokenRes{
			Token:        token,
			RefreshToken: refreshToken,
		},
	}

	return

}
