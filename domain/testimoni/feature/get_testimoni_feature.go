package feature

import (
	"context"
	"zayyid-go/domain/testimoni/model"
)

func (f TestimoniFeature) GetTestimoniFeature(ctx context.Context, request model.Testimoni) (response model.Testimoni, err error) {

	response, err = f.repo.GetTestimoniRepository(ctx, request)
	if err != nil {
		return
	}

	return
}
