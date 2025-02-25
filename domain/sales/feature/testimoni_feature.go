package feature

import (
	"context"
	paginate "zayyid-go/domain/shared/helper/pagination"
	sharedModel "zayyid-go/domain/shared/model"
	sharedRepo "zayyid-go/domain/shared/repository"

	modelRequest "zayyid-go/domain/sales/model/request"
)

func (f SalesFeature) AddTestimoniFeature(ctx context.Context, request modelRequest.Testimoni) (err error) {

	request.Id = sharedRepo.GenerateUuidAsIdTable().String()
	err = f.repo.AddTestimoniRepository(ctx, request)

	return
}

func (f SalesFeature) UpdateTestimoniFeature(ctx context.Context, request modelRequest.Testimoni) (err error) {

	err = f.repo.UpdateTestimoniRepository(ctx, request)

	return
}

func (f SalesFeature) GetTestimoniFeature(ctx context.Context, request modelRequest.Testimoni) (response modelRequest.Testimoni, err error) {

	response, err = f.repo.GetTestimoniRepository(ctx, request)

	return
}

func (f SalesFeature) GetListTestimoniFeature(ctx context.Context, request modelRequest.Testimoni, filter modelRequest.TestimoniSearch) (response []modelRequest.Testimoni, pagination *sharedModel.Pagination, err error) {

	response, err = f.repo.GetListTestimoniRepository(ctx, request, filter)
	if err != nil {
		return
	}

	count, err := f.repo.CountListTestimoniRepository(ctx, request)
	if err != nil {
		return
	}

	pagination, err = paginate.CalculatePagination(ctx, filter.Limit, count)
	if err != nil {
		return
	}

	pagination.Page = filter.Page

	return
}

func (f SalesFeature) GetPublicListTestimoniFeature(ctx context.Context, subDomain, referral string, filter modelRequest.TestimoniSearch) (response []modelRequest.Testimoni, pagination *sharedModel.Pagination, err error) {

	//validation referral code
	if referral != "" {

	}

	response, err = f.repo.GetPublicListTestimoniRepository(ctx, subDomain, filter)
	if err != nil {
		return
	}

	paramCount := modelRequest.Testimoni{
		UserName: subDomain,
	}
	count, err := f.repo.CountListTestimoniRepository(ctx, paramCount)
	if err != nil {
		return
	}

	pagination, err = paginate.CalculatePagination(ctx, filter.Limit, count)
	if err != nil {
		return
	}

	pagination.Page = filter.Page

	return
}
