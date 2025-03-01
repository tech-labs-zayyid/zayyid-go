package feature

import (
	"context"
	"errors"
	"net/http"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	sharedContext "zayyid-go/domain/shared/context"
	sharedHelper "zayyid-go/domain/shared/helper"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
)

func (f salesFeature) AddBannerSales(ctx context.Context, param request.BannerReq) (err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	tx := f.repo.OpenTransaction(ctx)

	defer func() {
		if err != nil {
			errRollback := f.repo.RollbackTransaction(tx)
			if errRollback != nil {
				err = sharedError.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), errRollback)
			}
		} else {
			errCommit := f.repo.CommitTransaction(tx)
			if errCommit != nil {
				err = sharedError.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), errCommit)
			}
		}
	}()

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"
	valueCtx.Username = "ekotoyota"

	//validation exists or not sales id

	for _, req := range param.DataBanner {
		err = sharedHelper.Validate(req)
		if err != nil {
			return
		}
	}

	count, err := f.repo.CountBannerSales(ctx, valueCtx.SalesId)
	if err != nil {
		return
	}

	if len(param.DataBanner) > 5 || count+len(param.DataBanner) > 5 {
		return sharedError.New(http.StatusBadRequest, sharedConstant.ErrMaximumUploadBanner, errors.New(sharedConstant.ErrMaximumUploadGallery))
	}

	param.SalesId = valueCtx.SalesId
	param.PublicAccess = valueCtx.Username
	err = f.repo.AddBannerSales(ctx, tx, param)
	return
}

func (f salesFeature) GetListBannerSales(ctx context.Context) (resp response.BannerListSalesResp, err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"

	//validation sales id

	resp, err = f.repo.GetListBannerSales(ctx, valueCtx.SalesId)
	return
}

func (f salesFeature) GetListBannerPublic(ctx context.Context, subdomain, referral string) (resp response.BannerListPublicSalesResp, err error) {
	//validation subdomain

	//validation referal
	if referral != "" {

	}

	resp, err = f.repo.GetListBannerPublicSales(ctx, subdomain)
	return
}

func (f salesFeature) GetBannerSales(ctx context.Context, id string) (resp response.BannerResp, err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"

	//validation sales id

	resp, err = f.repo.GetBannerSales(ctx, id, valueCtx.SalesId)
	return
}

func (f salesFeature) UpdateBanner(ctx context.Context, req request.BannerUpdateReq) (err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"

	err = sharedHelper.Validate(req)
	if err != nil {
		return
	}

	//validation sales id

	req.SalesId = valueCtx.SalesId
	err = f.repo.UpdateBannerSales(ctx, req)
	return
}
