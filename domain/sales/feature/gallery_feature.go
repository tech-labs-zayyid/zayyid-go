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

	count, err := f.repo.GetCountDataGalleryBySalesId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if len(param.ImageUrl) > 20 || count+len(param.ImageUrl) > 20 {
		return sharedError.New(http.StatusBadRequest, sharedConstant.ErrMaximumUploadGallery, errors.New(sharedConstant.ErrMaximumUploadGallery))
	}

	param.SalesId = valueCtx.UserId
	param.PublicAccess = dataUser.Username
	if err = f.repo.AddGallerySales(ctx, tx, param); err != nil {
		return
	}

	return
}

func (f salesFeature) GetDataListGallery(ctx context.Context) (resp response.GalleryResp, err error) {
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

	resp, err = f.repo.GetListDataGallerySales(ctx, valueCtx.UserId)
	return
}

func (f salesFeature) GetDataListGalleryPublic(ctx context.Context, subdomain string) (resp response.GalleryPublicResp, err error) {
	exists, err := f.userRepo.CheckExistsSubdomain(ctx, subdomain)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	resp, err = f.repo.GetListDataGalleryPublic(ctx, subdomain)
	return
}

func (f salesFeature) GetDataGallerySales(ctx context.Context, id string) (resp response.GalleryDataResp, err error) {
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

	existGalleryId, err := f.repo.CheckExistsGalleryId(ctx, id, valueCtx.UserId)
	if err != nil {
		return
	}

	if !existGalleryId {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrIdGalleryNotFound, errors.New(sharedConstant.ErrIdGalleryNotFound))
		return
	}

	resp, err = f.repo.GetDataGallerySales(ctx, id, valueCtx.UserId)
	return
}

func (f salesFeature) UpdateGallery(ctx context.Context, req request.UpdateGalleryParam) (err error) {
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

	err = sharedHelper.Validate(req)
	if err != nil {
		return
	}

	existGalleryId, err := f.repo.CheckExistsGalleryId(ctx, req.Id, valueCtx.UserId)
	if err != nil {
		return
	}

	if !existGalleryId {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrIdGalleryNotFound, errors.New(sharedConstant.ErrIdGalleryNotFound))
		return
	}

	req.SalesId = valueCtx.SalesId
	err = f.repo.UpdateGallerySales(ctx, req)
	return
}
