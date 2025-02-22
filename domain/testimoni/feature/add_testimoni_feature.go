package feature

import (
	"context"
	sharedRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/domain/testimoni/model"
)

func (f TestimoniFeature) AddTestimoniFeature(ctx context.Context, request model.Testimoni) (err error) {

	request.Id = sharedRepo.GenerateUuidAsIdTable().String()
	err = f.repo.AddTestimoniRepository(ctx, request)
	if err != nil {
		return
	}

	return
}
