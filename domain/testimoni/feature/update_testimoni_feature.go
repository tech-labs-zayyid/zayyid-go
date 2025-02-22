package feature

import (
	"context"
	"zayyid-go/domain/testimoni/model"
)

func (f TestimoniFeature) UpdateTestimoniFeature(ctx context.Context, request model.Testimoni) (err error) {

	err = f.repo.UpdateTestimoniRepository(ctx, request)
	if err != nil {
		return
	}

	return
}
