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

func (f salesFeature) AddGallerySales(ctx context.Context, param request.AddGalleryParam) (err error) {
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

	//validation exists or not sales id in t_gallery

	count, err := f.repo.GetCountDataGalleryBySalesId(ctx, valueCtx.SalesId)
	if err != nil {
		return
	}

	if len(param.ImageUrl) > 20 || count+len(param.ImageUrl) > 20 {
		return sharedError.New(http.StatusBadRequest, sharedConstant.ErrMaximumUploadGallery, errors.New(sharedConstant.ErrMaximumUploadGallery))
	}

	param.SalesId = valueCtx.SalesId
	param.PublicAccess = valueCtx.Username
	if err = f.repo.AddGallerySales(ctx, tx, param); err != nil {
		return
	}

	return
}

func (f salesFeature) GetDataListGallery(ctx context.Context) (resp response.GalleryResp, err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"

	//validation sales id

	resp, err = f.repo.GetListDataGallerySales(ctx, valueCtx.SalesId)
	return
}

func (f salesFeature) GetDataListGalleryPublic(ctx context.Context, subdomain, referral string) (resp response.GalleryPublicResp, err error) {
	//validation subdomain

	//validation referral code
	if referral != "" {

	}

	resp, err = f.repo.GetListDataGalleryPublic(ctx, subdomain)
	return
}

func (f salesFeature) GetDataGallerySales(ctx context.Context, id string) (resp response.GalleryDataResp, err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"

	//validation sales id

	resp, err = f.repo.GetDataGallerySales(ctx, id, valueCtx.SalesId)
	return
}

func (f salesFeature) UpdateGallery(ctx context.Context, req request.UpdateGalleryParam) (err error) {
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
	err = f.repo.UpdateGallerySales(ctx, req)
	return
}
