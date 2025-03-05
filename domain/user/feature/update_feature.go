package feature

import (
	"context"
	sharedHelper "zayyid-go/domain/shared/helper/general"
	"zayyid-go/domain/user/model"
)

func (f UserFeature) UpdateFeature(ctx context.Context, payload model.UpdateUser, userId string) (err error) {

	// hash password when not null
	if payload.Password != "" {
		encryptedPassword, errHashPassword := sharedHelper.HashPassword(payload.Password)
		if errHashPassword != nil {
			err = errHashPassword
			return
		}

		payload.Password = encryptedPassword
	}

	// update data
	err = f.repo.UpdateRepository(ctx, payload, userId)
	if err != nil {
		return
	}

	return

}
