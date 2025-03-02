package feature

import (
	"context"
	"zayyid-go/domain/user/model"
)

func (f UserFeature) UpdateFeature(ctx context.Context, payload model.UpdateUser, userId string) (err error) {

	// update data
	err = f.repo.UpdateRepository(ctx, payload, userId)
	if err != nil {
		return
	}

	return

}
