package feature

import (
	"context"
	"errors"
	"net/http"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	sharedContext "zayyid-go/domain/shared/context"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedHelper "zayyid-go/domain/shared/helper/general"
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

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	dataUser, err := f.userRepo.GetDataUserByUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

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
	param.PublicAccess = dataUser.Username
	err = f.repo.AddBannerSales(ctx, tx, param)
	return
}

func (f salesFeature) GetListBannerSales(ctx context.Context) (resp response.BannerListSalesResp, err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	resp, err = f.repo.GetListBannerSales(ctx, valueCtx.UserId)
	return
}

func (f salesFeature) GetListBannerPublic(ctx context.Context, subdomain string) (resp response.BannerListPublicSalesResp, err error) {
	exists, err := f.userRepo.CheckExistsSubdomain(ctx, subdomain)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusNotFound, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	resp, err = f.repo.GetListBannerPublicSales(ctx, subdomain)
	return
}

func (f salesFeature) GetBannerSales(ctx context.Context, id string) (resp response.BannerResp, err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	existBannerId, err := f.repo.CheckExistsBannerId(ctx, id, valueCtx.UserId)
	if err != nil {
		return
	}

	if !existBannerId {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrIdBannerNotFound, errors.New(sharedConstant.ErrIdBannerNotFound))
		return
	}

	resp, err = f.repo.GetBannerSales(ctx, id, valueCtx.SalesId)
	return
}

func (f salesFeature) UpdateBanner(ctx context.Context, req request.BannerUpdateReq) (err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	existBannerId, err := f.repo.CheckExistsBannerId(ctx, req.Id, valueCtx.UserId)
	if err != nil {
		return
	}

	if !existBannerId {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrIdBannerNotFound, errors.New(sharedConstant.ErrIdBannerNotFound))
		return
	}

	err = sharedHelper.Validate(req)
	if err != nil {
		return
	}

	req.SalesId = valueCtx.UserId
	err = f.repo.UpdateBannerSales(ctx, req)
	return
}
