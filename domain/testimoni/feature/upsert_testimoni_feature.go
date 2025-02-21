package feature

import (
	"context"
	sharedRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/domain/testimoni/model"
)

func (f TestimoniFeature) UpsertTestimoniFeature(ctx context.Context, request model.Testimoni) (err error) {
	if request.IsUpdate == 0 {
		request.Id = sharedRepo.GenerateUuidAsIdTable().String()
		err = f.repo.AddTestimoniRepository(ctx, request)
	} else {
		err = f.repo.UpdateTestimoniRepository(ctx, request)
	}

	if err != nil {
		return
	}

	return
}
