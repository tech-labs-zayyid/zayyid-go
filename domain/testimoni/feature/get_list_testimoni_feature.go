package feature

import (
	"context"
	paginate "zayyid-go/domain/shared/helper/pagination"
	sharedModel "zayyid-go/domain/shared/model"
	"zayyid-go/domain/testimoni/model"
)

func (f TestimoniFeature) GetListTestimoniFeature(ctx context.Context, request model.Testimoni) (response []model.Testimoni, pagination *sharedModel.Pagination, err error) {

	response, err = f.repo.GetListTestimoniRepository(ctx, request)
	if err != nil {
		return
	}

	count, err := f.repo.CountListTestimoniRepository(ctx, request)
	if err != nil {
		return
	}

	pagination, err = paginate.CalculatePagination(ctx, request.Limit, count)
	if err != nil {
		return
	}

	pagination.Page = request.Page

	return
}
